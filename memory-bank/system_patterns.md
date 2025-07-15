# System Patterns - Chess Tournament Hub

## System Architecture Overview

### High-Level Architecture
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │   Mobile App    │    │   Admin Panel   │
│   (Next.js)     │    │ (React Native)  │    │   (Next.js)     │
└─────────┬───────┘    └─────────┬───────┘    └─────────┬───────┘
          │                      │                      │
          └──────────────────────┼──────────────────────┘
                                 │
                    ┌────────────▼────────────┐
                    │      Load Balancer      │
                    │       (AWS ALB)         │
                    └────────────┬────────────┘
                                 │
                    ┌────────────▼────────────┐
                    │      API Gateway       │
                    │    (Go + Gin Web)      │
                    └────────────┬────────────┘
                                 │
          ┌──────────────────────┼──────────────────────┐
          │                      │                      │
  ┌───────▼───────┐    ┌─────────▼─────────┐    ┌───────▼───────┐
  │  Auth Service │    │Tournament Service │    │Payment Service│
  │      (Go)     │    │       (Go)        │    │     (Go)      │
  └───────┬───────┘    └─────────┬─────────┘    └───────┬───────┘
          │                      │                      │
          └──────────────────────┼──────────────────────┘
                                 │
                    ┌────────────▼────────────┐
                    │     Database Layer      │
                    │    PostgreSQL + Redis   │
                    └─────────────────────────┘
```

### Service Architecture Pattern

#### Modular Monolith Approach
- **Single Deployment Unit:** One Go application with modular internal structure
- **Service Boundaries:** Clear separation between auth, tournaments, payments, notifications
- **Shared Database:** Single PostgreSQL instance with logical separation
- **Future Migration Path:** Can extract to microservices when scale demands

#### Service Organization
```go
internal/
├── auth/           // Authentication and authorization
├── users/          // User profile management
├── tournaments/    // Tournament CRUD and management
├── registrations/  // Player registration handling
├── payments/       // Payment processing and escrow
├── notifications/  // Email and SMS notifications
├── analytics/      // Usage analytics and reporting
└── shared/         // Common utilities and middleware
```

## Key Technical Decisions

### Database Design Patterns

#### Single Database with Logical Separation
```sql
-- User Management
users (id, email, role, profile_data, created_at)
player_profiles (user_id, chess_rating, preferences)
organizer_profiles (user_id, organization_data, verification_status)

-- Tournament Management
tournaments (id, organizer_id, tournament_data, status, created_at)
registrations (id, tournament_id, player_id, payment_status, registered_at)
tournament_rounds (id, tournament_id, round_number, pairings, results)

-- Payment Management
payments (id, tournament_id, player_id, amount, stripe_payment_id, status)
escrow_transactions (id, payment_id, release_date, organizer_payout_id)
organizer_payouts (id, organizer_id, total_amount, stripe_transfer_id)
```

#### JSON Column Pattern for Flexibility
```go
type Tournament struct {
    ID           string          `json:"id" db:"id"`
    OrganizerID  string          `json:"organizer_id" db:"organizer_id"`
    Title        string          `json:"title" db:"title"`
    Config       TournamentConfig `json:"config" db:"config"` // JSON column
    Status       string          `json:"status" db:"status"`
    CreatedAt    time.Time       `json:"created_at" db:"created_at"`
}

type TournamentConfig struct {
    Format       string      `json:"format"`
    TimeControl  TimeControl `json:"time_control"`
    Venue        Venue       `json:"venue"`
    Requirements Requirements `json:"requirements"`
    PrizePool    []Prize     `json:"prize_pool"`
}
```

### API Design Patterns

#### RESTful API with Consistent Structure
```go
// Standard endpoint patterns
GET    /api/v1/tournaments              // List tournaments
POST   /api/v1/tournaments              // Create tournament
GET    /api/v1/tournaments/{id}         // Get tournament details
PUT    /api/v1/tournaments/{id}         // Update tournament
DELETE /api/v1/tournaments/{id}         // Cancel tournament

// Nested resource patterns
POST   /api/v1/tournaments/{id}/register    // Register for tournament
GET    /api/v1/tournaments/{id}/participants // Get participants
POST   /api/v1/tournaments/{id}/rounds      // Create tournament round
```

#### Standardized Response Format
```go
type APIResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   *APIError   `json:"error,omitempty"`
    Meta    *MetaData   `json:"meta,omitempty"`
}

type APIError struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Details string `json:"details,omitempty"`
}

type MetaData struct {
    Page       int `json:"page,omitempty"`
    Limit      int `json:"limit,omitempty"`
    Total      int `json:"total,omitempty"`
    TotalPages int `json:"total_pages,omitempty"`
}
```

### Authentication & Authorization Patterns

#### JWT-Based Authentication
```go
type AuthService struct {
    jwtSecret     []byte
    tokenExpiry   time.Duration
    refreshExpiry time.Duration
}

type JWTClaims struct {
    UserID   string   `json:"user_id"`
    Email    string   `json:"email"`
    Role     string   `json:"role"`
    Permissions []string `json:"permissions"`
    jwt.StandardClaims
}

// Middleware pattern for route protection
func (a *AuthService) RequireAuth() gin.HandlerFunc {
    return gin.HandlerFunc(func(c *gin.Context) {
        token := extractTokenFromHeader(c)
        claims, err := a.ValidateToken(token)
        if err != nil {
            c.JSON(401, APIResponse{Error: &APIError{Code: "UNAUTHORIZED"}})
            c.Abort()
            return
        }
        c.Set("user", claims)
        c.Next()
    })
}
```

#### Role-Based Access Control
```go
type Permission string

const (
    PermissionCreateTournament Permission = "tournament:create"
    PermissionManageTournament Permission = "tournament:manage"
    PermissionRegisterPlayer   Permission = "player:register"
    PermissionViewAnalytics    Permission = "analytics:view"
)

func (a *AuthService) RequirePermission(perm Permission) gin.HandlerFunc {
    return gin.HandlerFunc(func(c *gin.Context) {
        user := c.MustGet("user").(*JWTClaims)
        if !hasPermission(user.Permissions, string(perm)) {
            c.JSON(403, APIResponse{Error: &APIError{Code: "FORBIDDEN"}})
            c.Abort()
            return
        }
        c.Next()
    })
}
```

## Design Patterns in Use

### Repository Pattern for Data Access
```go
type TournamentRepository interface {
    Create(ctx context.Context, tournament *Tournament) error
    GetByID(ctx context.Context, id string) (*Tournament, error)
    List(ctx context.Context, filters TournamentFilters) ([]*Tournament, error)
    Update(ctx context.Context, tournament *Tournament) error
    Delete(ctx context.Context, id string) error
}

type tournamentRepository struct {
    db *sql.DB
}

func (r *tournamentRepository) Create(ctx context.Context, tournament *Tournament) error {
    query := `INSERT INTO tournaments (id, organizer_id, title, config, status, created_at) 
              VALUES ($1, $2, $3, $4, $5, $6)`
    _, err := r.db.ExecContext(ctx, query, 
        tournament.ID, tournament.OrganizerID, tournament.Title, 
        tournament.Config, tournament.Status, tournament.CreatedAt)
    return err
}
```

### Service Layer Pattern
```go
type TournamentService struct {
    repo         TournamentRepository
    paymentSvc   PaymentService
    notificationSvc NotificationService
    logger       *logrus.Logger
}

func (s *TournamentService) CreateTournament(ctx context.Context, req CreateTournamentRequest) (*Tournament, error) {
    // Business logic validation
    if err := s.validateTournamentRequest(req); err != nil {
        return nil, err
    }

    // Create tournament entity
    tournament := &Tournament{
        ID:          generateID(),
        OrganizerID: req.OrganizerID,
        Title:       req.Title,
        Config:      req.Config,
        Status:      "draft",
        CreatedAt:   time.Now(),
    }

    // Persist to database
    if err := s.repo.Create(ctx, tournament); err != nil {
        s.logger.WithError(err).Error("Failed to create tournament")
        return nil, err
    }

    // Send notifications
    go s.notificationSvc.SendTournamentCreated(tournament)

    return tournament, nil
}
```

### Event-Driven Pattern for Decoupling
```go
type Event struct {
    Type      string      `json:"type"`
    ID        string      `json:"id"`
    Timestamp time.Time   `json:"timestamp"`
    Data      interface{} `json:"data"`
}

type EventBus interface {
    Publish(event Event) error
    Subscribe(eventType string, handler EventHandler) error
}

type EventHandler func(event Event) error

// Usage in service
func (s *TournamentService) PublishTournament(ctx context.Context, tournamentID string) error {
    tournament, err := s.repo.GetByID(ctx, tournamentID)
    if err != nil {
        return err
    }

    tournament.Status = "published"
    if err := s.repo.Update(ctx, tournament); err != nil {
        return err
    }

    // Publish event for other services to react
    event := Event{
        Type:      "tournament.published",
        ID:        generateEventID(),
        Timestamp: time.Now(),
        Data:      tournament,
    }

    return s.eventBus.Publish(event)
}
```

## Component Relationships

### Frontend Component Architecture

#### React Component Hierarchy
```
App
├── AuthProvider (Context for authentication state)
├── Router
│   ├── PublicRoutes
│   │   ├── LandingPage
│   │   ├── TournamentListing
│   │   ├── TournamentDetail
│   │   └── AuthPages (Login/Register)
│   └── ProtectedRoutes
│       ├── Dashboard (Player/Organizer specific)
│       ├── TournamentCreation
│       ├── TournamentManagement
│       ├── Profile
│       └── PaymentHistory
```

#### State Management Pattern
```typescript
// Context-based state for global app state
interface AppContextType {
  user: User | null;
  tournaments: Tournament[];
  loading: boolean;
  error: string | null;
}

// React Query for server state
const useTournaments = (filters: TournamentFilters) => {
  return useQuery({
    queryKey: ['tournaments', filters],
    queryFn: () => tournamentApi.list(filters),
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
};

// Local state for component-specific state
const [formData, setFormData] = useState<CreateTournamentForm>({
  title: '',
  description: '',
  startDate: '',
  // ...
});
```

#### Component Communication Patterns
```typescript
// Props drilling for simple parent-child communication
<TournamentCard 
  tournament={tournament}
  onRegister={handleRegister}
  isRegistered={isUserRegistered}
/>

// Context for deeply nested components
const { user, updateUser } = useAuth();

// Event emitters for loosely coupled components
const eventBus = new EventTarget();
eventBus.dispatchEvent(new CustomEvent('tournament:registered', {
  detail: { tournamentId, userId }
}));
```

## Critical Implementation Paths

### User Registration and Authentication Flow
```go
// 1. User Registration
POST /api/v1/auth/register
→ Validate input data
→ Hash password with bcrypt
→ Create user record with pending status
→ Generate email verification token
→ Send verification email
→ Return success response

// 2. Email Verification
GET /api/v1/auth/verify?token=xyz
→ Validate token and expiry
→ Update user status to active
→ Generate JWT tokens
→ Redirect to dashboard

// 3. Login Process
POST /api/v1/auth/login
→ Validate credentials
→ Generate access and refresh tokens
→ Update last login timestamp
→ Return tokens and user profile
```

### Tournament Creation and Registration Flow
```go
// 1. Tournament Creation
POST /api/v1/tournaments
→ Validate organizer permissions
→ Validate tournament data
→ Create tournament in draft status
→ Generate unique tournament URL
→ Send confirmation to organizer

// 2. Tournament Publication
PUT /api/v1/tournaments/{id}/publish
→ Validate tournament completeness
→ Update status to published
→ Index for search
→ Send notifications to followers

// 3. Player Registration
POST /api/v1/tournaments/{id}/register
→ Check tournament capacity
→ Validate player eligibility
→ Create registration record
→ Initiate payment process
→ Send confirmation email
```

### Payment Processing Flow
```go
// 1. Payment Initiation
POST /api/v1/payments/create-intent
→ Calculate total amount (entry fee + platform fee)
→ Create Stripe Payment Intent
→ Store payment record with pending status
→ Return client secret for frontend

// 2. Payment Confirmation
POST /api/v1/payments/confirm
→ Confirm payment with Stripe
→ Update registration status
→ Add funds to escrow
→ Send payment confirmation

// 3. Escrow Release (automated)
CRON: Process escrow releases
→ Find tournaments completed > 7 days ago
→ Calculate organizer payout (total - commission)
→ Transfer to organizer's Stripe account
→ Update payout records
→ Send payout confirmation
```

### Tournament Execution Flow
```go
// 1. Tournament Start
PUT /api/v1/tournaments/{id}/start
→ Validate all participants checked in
→ Generate initial pairings
→ Create first round
→ Notify participants of pairings

// 2. Round Management
POST /api/v1/tournaments/{id}/rounds/{round}/results
→ Validate arbiter permissions
→ Update match results
→ Calculate new standings
→ Generate next round pairings (if applicable)
→ Notify participants

// 3. Tournament Completion
PUT /api/v1/tournaments/{id}/complete
→ Finalize all results
→ Calculate final standings
→ Update player ratings
→ Trigger escrow release
→ Generate tournament report
```

## Data Flow Patterns

### Request-Response Pattern
```
Client Request → Middleware (Auth, Logging) → Router → Handler → Service → Repository → Database
                    ↓
Client Response ← JSON Serialization ← Service Response ← Database Result
```

### Event-Driven Data Flow
```
Action Trigger → Event Publication → Event Queue → Event Handlers → Side Effects
                                                     ├── Email Notifications
                                                     ├── Rating Updates
                                                     ├── Analytics Tracking
                                                     └── Cache Invalidation
```

### Caching Strategy
```go
// Multi-level caching approach
type CacheService struct {
    redis    *redis.Client
    memory   *bigcache.BigCache
    postgres *sql.DB
}

// Cache hierarchy: Memory → Redis → Database
func (c *CacheService) GetTournament(id string) (*Tournament, error) {
    // L1 Cache: In-memory (fastest)
    if data, err := c.memory.Get(id); err == nil {
        return unmarshalTournament(data), nil
    }

    // L2 Cache: Redis (fast)
    if data, err := c.redis.Get(ctx, id).Result(); err == nil {
        tournament := unmarshalTournament(data)
        c.memory.Set(id, data) // Populate L1
        return tournament, nil
    }

    // L3 Source: Database (slowest)
    tournament, err := c.postgres.GetTournament(id)
    if err != nil {
        return nil, err
    }

    // Populate caches
    data := marshalTournament(tournament)
    c.redis.Set(ctx, id, data, time.Hour)
    c.memory.Set(id, data)

    return tournament, nil
}
```

## Error Handling Patterns

### Centralized Error Handling
```go
type ErrorCode string

const (
    ErrInvalidInput     ErrorCode = "INVALID_INPUT"
    ErrUnauthorized     ErrorCode = "UNAUTHORIZED"
    ErrTournamentFull   ErrorCode = "TOURNAMENT_FULL"
    ErrPaymentFailed    ErrorCode = "PAYMENT_FAILED"
)

type AppError struct {
    Code    ErrorCode `json:"code"`
    Message string    `json:"message"`
    Details string    `json:"details,omitempty"`
    Cause   error     `json:"-"`
}

func (e *AppError) Error() string {
    return e.Message
}

// Middleware for error handling
func ErrorHandlerMiddleware() gin.HandlerFunc {
    return gin.HandlerFunc(func(c *gin.Context) {
        c.Next()

        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err
            
            var appErr *AppError
            if errors.As(err, &appErr) {
                c.JSON(getHTTPStatus(appErr.Code), APIResponse{
                    Success: false,
                    Error:   appErr,
                })
            } else {
                c.JSON(500, APIResponse{
                    Success: false,
                    Error: &AppError{
                        Code:    "INTERNAL_ERROR",
                        Message: "An unexpected error occurred",
                    },
                })
            }
        }
    })
}
```

### Graceful Degradation Pattern
```go
// Circuit breaker for external services
type CircuitBreaker struct {
    maxFailures int
    timeout     time.Duration
    failures    int
    lastFailure time.Time
    state       CBState
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    if cb.state == CBStateOpen {
        if time.Since(cb.lastFailure) > cb.timeout {
            cb.state = CBStateHalfOpen
        } else {
            return errors.New("circuit breaker open")
        }
    }

    err := fn()
    if err != nil {
        cb.failures++
        cb.lastFailure = time.Now()
        if cb.failures >= cb.maxFailures {
            cb.state = CBStateOpen
        }
        return err
    }

    cb.failures = 0
    cb.state = CBStateClosed
    return nil
}
```

## Performance Patterns

### Database Query Optimization
```go
// Pagination pattern
type PaginationParams struct {
    Page  int `json:"page" form:"page"`
    Limit int `json:"limit" form:"limit"`
}

func (r *tournamentRepository) ListWithPagination(
    ctx context.Context, 
    filters TournamentFilters, 
    pagination PaginationParams,
) ([]*Tournament, int, error) {
    offset := (pagination.Page - 1) * pagination.Limit
    
    query := `
        SELECT t.*, COUNT(*) OVER() as total_count
        FROM tournaments t
        WHERE status = $1
        ORDER BY start_date ASC
        LIMIT $2 OFFSET $3
    `
    
    rows, err := r.db.QueryContext(ctx, query, "published", pagination.Limit, offset)
    // ... handle results
}

// Index optimization
CREATE INDEX CONCURRENTLY idx_tournaments_status_start_date 
ON tournaments(status, start_date) 
WHERE status IN ('published', 'ongoing');

CREATE INDEX CONCURRENTLY idx_registrations_tournament_player 
ON registrations(tournament_id, player_id);
```

### Connection Pooling Pattern
```go
// Database connection pool configuration
func NewDB(databaseURL string) (*sql.DB, error) {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return nil, err
    }

    // Connection pool settings
    db.SetMaxOpenConns(25)                 // Maximum connections
    db.SetMaxIdleConns(5)                  // Idle connections
    db.SetConnMaxLifetime(5 * time.Minute) // Connection lifetime
    db.SetConnMaxIdleTime(1 * time.Minute) // Idle timeout

    return db, nil
}
```

## Security Patterns

### Input Validation and Sanitization
```go
type TournamentCreateRequest struct {
    Title       string    `json:"title" validate:"required,min=5,max=100"`
    Description string    `json:"description" validate:"max=1000"`
    StartDate   time.Time `json:"start_date" validate:"required,future"`
    EntryFee    float64   `json:"entry_fee" validate:"min=0,max=1000"`
    MaxPlayers  int       `json:"max_players" validate:"required,min=4,max=500"`
}

func validateAndSanitize(req *TournamentCreateRequest) error {
    // Validate using struct tags
    if err := validator.New().Struct(req); err != nil {
        return err
    }

    // Sanitize HTML content
    req.Description = bluemonday.UGCPolicy().Sanitize(req.Description)
    req.Title = strings.TrimSpace(req.Title)

    return nil
}
```

### Rate Limiting Pattern
```go
// Redis-based rate limiting
type RateLimiter struct {
    redis  *redis.Client
    window time.Duration
    limit  int
}

func (rl *RateLimiter) Allow(key string) (bool, error) {
    current := time.Now().Unix()
    window := rl.window.Seconds()
    
    pipe := rl.redis.Pipeline()
    pipe.ZRemRangeByScore(ctx, key, "-inf", fmt.Sprintf("%d", current-int64(window)))
    pipe.ZCard(ctx, key)
    pipe.ZAdd(ctx, key, &redis.Z{Score: float64(current), Member: current})
    pipe.Expire(ctx, key, rl.window)
    
    results, err := pipe.Exec(ctx)
    if err != nil {
        return false, err
    }
    
    count := results[1].(*redis.IntCmd).Val()
    return count < int64(rl.limit), nil
}
```

This system patterns document establishes the foundational architecture and design patterns that will guide all technical implementation decisions throughout the project lifecycle.
