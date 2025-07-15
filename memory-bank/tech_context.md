# Tech Context - Chess Tournament Hub

## Technology Stack

### Backend Technologies

#### Core Framework
- **Language:** Go 1.21+
- **Web Framework:** Gin (high-performance HTTP web framework)
- **Reason:** Superior concurrency handling for tournament registrations, excellent performance, strong typing

#### Database Stack
- **Primary Database:** PostgreSQL 15+
  - ACID compliance for financial transactions
  - JSON column support for flexible tournament configurations
  - Excellent performance with proper indexing
  - Strong ecosystem and tooling
- **Caching Layer:** Redis 7+
  - Session management and JWT token blacklisting
  - Tournament listing cache
  - Rate limiting and request throttling
  - Real-time data for live tournament updates

#### Authentication & Security
- **JWT Tokens:** golang-jwt/jwt for stateless authentication
- **Password Hashing:** bcrypt for secure password storage
- **Input Validation:** go-playground/validator for request validation
- **Security Headers:** Gin middleware for CORS, security headers
- **Rate Limiting:** Redis-based sliding window rate limiting

#### External Integrations
- **Payment Processing:** Stripe Connect API
  - Malaysian payment methods (FPX, credit cards)
  - Escrow and marketplace functionality
  - Automated payouts and commission handling
- **Email Service:** SendGrid or AWS SES
  - Transactional emails for confirmations
  - Bulk email for tournament announcements
- **SMS Service:** Twilio or local Malaysian SMS provider
  - Phone verification and tournament reminders

### Frontend Technologies

#### Core Framework
- **Framework:** Next.js 14+ with App Router
- **Language:** TypeScript for type safety
- **Styling:** Tailwind CSS for utility-first styling
- **Reason:** Server-side rendering for SEO, excellent developer experience, strong ecosystem

#### State Management
- **Server State:** TanStack Query (React Query)
  - Caching and synchronization of server data
  - Background updates and error handling
  - Optimistic updates for better UX
- **Client State:** React Context + useState/useReducer
  - Authentication state
  - UI state and form data
  - Shopping cart equivalent for tournament registration

#### UI Components
- **Component Library:** Custom components built with Tailwind
- **Icons:** Lucide React for consistent iconography
- **Forms:** React Hook Form with Zod validation
- **Charts:** Recharts for tournament analytics
- **Date/Time:** date-fns for date manipulation

#### Development Tools
- **Linting:** ESLint with TypeScript rules
- **Formatting:** Prettier for code formatting
- **Testing:** Jest + React Testing Library
- **Build Tool:** Next.js built-in bundling

### Mobile Technologies

#### Framework (Phase 2)
- **Framework:** React Native with Expo
- **Language:** TypeScript
- **Navigation:** React Navigation
- **State Management:** Same as web (TanStack Query + Context)
- **Reason:** Code sharing with web frontend, rapid development, excellent ecosystem

### Infrastructure Technologies

#### Cloud Platform
- **Primary:** AWS (Amazon Web Services)
- **Compute:** ECS with Fargate for containerized Go backend
- **Database:** RDS PostgreSQL with Multi-AZ deployment
- **Caching:** ElastiCache Redis cluster
- **Storage:** S3 for file uploads (tournament documents, logos)
- **CDN:** CloudFront for static asset delivery

#### Development & Deployment
- **Containerization:** Docker for local development and production
- **Orchestration:** AWS ECS for container orchestration
- **Load Balancing:** Application Load Balancer with health checks
- **Auto Scaling:** ECS auto-scaling based on CPU/memory metrics

#### Monitoring & Observability
- **Application Monitoring:** AWS CloudWatch with custom metrics
- **Error Tracking:** Sentry for error reporting and performance monitoring
- **Logging:** Structured logging with Logrus, centralized in CloudWatch
- **Uptime Monitoring:** External service for health checks

## Development Setup

### Local Development Environment

#### Backend Setup
```bash
# Prerequisites
go version # 1.21+
docker --version
docker-compose --version

# Project setup
git clone <backend-repo>
cd chess-tournaments-api
cp .env.example .env

# Database setup
docker-compose up -d postgres redis

# Install dependencies and run
go mod download
go run cmd/server/main.go

# Database migrations
migrate -path migrations -database postgres://... up
```

#### Frontend Setup
```bash
# Prerequisites
node --version # 18+
npm --version # 9+

# Project setup
git clone <frontend-repo>
cd chess-tournaments-web
cp .env.local.example .env.local
npm install

# Development server
npm run dev

# Type checking and linting
npm run type-check
npm run lint
```

#### Full Stack Development
```bash
# Using Docker Compose for full stack
git clone <project-repo>
cd chess-tournaments
docker-compose up -d

# Services available at:
# Backend API: http://localhost:8080
# Frontend: http://localhost:3000
# PostgreSQL: localhost:5432
# Redis: localhost:6379
```

### Development Tools Configuration

#### VSCode Extensions
```json
{
  "recommendations": [
    "golang.go",
    "bradlc.vscode-tailwindcss",
    "esbenp.prettier-vscode",
    "ms-vscode.vscode-typescript-next",
    "ms-vscode-remote.remote-containers"
  ]
}
```

#### Git Hooks (pre-commit)
```yaml
# .pre-commit-config.yaml
repos:
  - repo: local
    hooks:
      - id: go-fmt
        name: go fmt
        entry: gofmt -w
        language: system
        files: \.go$
      - id: go-lint
        name: golangci-lint
        entry: golangci-lint run
        language: system
        files: \.go$
      - id: prettier
        name: prettier
        entry: prettier --write
        language: node
        files: \.(ts|tsx|js|jsx|json|css|md)$
```

## Technical Constraints

### Performance Requirements
- **API Response Time:** < 200ms for 95% of requests
- **Page Load Time:** < 2 seconds for tournament listing
- **Database Query Time:** < 100ms for complex tournament queries
- **Concurrent Users:** Support 1,000+ simultaneous users
- **File Upload Size:** Max 10MB for tournament documents

### Security Requirements
- **Data Encryption:** TLS 1.3 for all communications
- **Password Security:** Minimum 8 characters, bcrypt hashing
- **Session Management:** JWT with 15-minute access tokens, 7-day refresh tokens
- **Input Validation:** All user inputs validated and sanitized
- **Rate Limiting:** 100 requests per minute per IP for public endpoints

### Compliance Requirements
- **PCI DSS:** Level 4 compliance for payment processing
- **GDPR-style Privacy:** Malaysian PDPA compliance
- **Data Retention:** User data retention policies and deletion capabilities
- **Audit Logging:** Comprehensive logging for financial transactions

### Scalability Constraints
- **Database:** PostgreSQL with read replicas for scaling reads
- **Caching:** Redis cluster for distributed caching
- **File Storage:** S3 with CloudFront for global content delivery
- **Auto Scaling:** Horizontal scaling based on traffic patterns

## Dependencies

### Backend Dependencies
```go
// Core dependencies
github.com/gin-gonic/gin v1.9.1
github.com/lib/pq v1.10.9 // PostgreSQL driver
github.com/go-redis/redis/v8 v8.11.5
github.com/golang-jwt/jwt/v5 v5.0.0

// Utility dependencies
github.com/go-playground/validator/v10 v10.15.5
github.com/sirupsen/logrus v1.9.3
github.com/golang-migrate/migrate/v4 v4.16.2
github.com/joho/godotenv v1.4.0

// External service SDKs
github.com/stripe/stripe-go/v75 v75.9.0
github.com/sendgrid/sendgrid-go v3.12.0
github.com/aws/aws-sdk-go v1.44.327
```

### Frontend Dependencies
```json
{
  "dependencies": {
    "next": "^14.0.0",
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "@tanstack/react-query": "^5.0.0",
    "react-hook-form": "^7.45.0",
    "zod": "^3.22.0",
    "@stripe/stripe-js": "^2.1.0",
    "date-fns": "^2.30.0",
    "lucide-react": "^0.263.1"
  },
  "devDependencies": {
    "typescript": "^5.0.0",
    "tailwindcss": "^3.3.0",
    "eslint": "^8.45.0",
    "prettier": "^3.0.0",
    "@types/react": "^18.2.0"
  }
}
```

### Infrastructure Dependencies
```yaml
# docker-compose.yml
version: '3.8'
services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: chess_tournaments
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

  backend:
    build: ./backend
    ports:
      - "8080:8080"
    depends_on:
      - postgres
      - redis
    environment:
      DATABASE_URL: postgres://postgres:password@postgres:5432/chess_tournaments
      REDIS_URL: redis://redis:6379

  frontend:
    build: ./frontend
    ports:
      - "3000:3000"
    depends_on:
      - backend
    environment:
      NEXT_PUBLIC_API_URL: http://localhost:8080
```

## Tool Usage Patterns

### Database Migrations
```bash
# Create new migration
migrate create -ext sql -dir migrations add_tournaments_table

# Run migrations
migrate -path migrations -database $DATABASE_URL up

# Rollback migration
migrate -path migrations -database $DATABASE_URL down 1

# Check migration status
migrate -path migrations -database $DATABASE_URL version
```

### Testing Patterns
```go
// Backend testing
func TestCreateTournament(t *testing.T) {
    // Setup test database
    db := setupTestDB(t)
    defer teardownTestDB(t, db)
    
    service := NewTournamentService(db)
    
    // Test cases
    tournament, err := service.CreateTournament(context.Background(), validRequest)
    assert.NoError(t, err)
    assert.NotEmpty(t, tournament.ID)
}

// Run tests
go test ./... -v
go test -race ./... # Race condition detection
go test -cover ./... # Coverage report
```

```typescript
// Frontend testing
import { render, screen, fireEvent } from '@testing-library/react'
import { TournamentCard } from '../TournamentCard'

describe('TournamentCard', () => {
  it('renders tournament information', () => {
    render(<TournamentCard tournament={mockTournament} />)
    expect(screen.getByText('Test Tournament')).toBeInTheDocument()
    expect(screen.getByText('$50 entry fee')).toBeInTheDocument()
  })

  it('handles registration click', () => {
    const handleRegister = jest.fn()
    render(<TournamentCard tournament={mockTournament} onRegister={handleRegister} />)
    
    fireEvent.click(screen.getByText('Register'))
    expect(handleRegister).toHaveBeenCalledWith(mockTournament.id)
  })
})
```

### API Documentation
```go
// Swagger annotations in Go
// @Summary Create tournament
// @Description Create a new chess tournament
// @Tags tournaments
// @Accept json
// @Produce json
// @Param tournament body CreateTournamentRequest true "Tournament data"
// @Success 201 {object} Tournament
// @Failure 400 {object} APIError
// @Router /tournaments [post]
func (h *TournamentHandler) CreateTournament(c *gin.Context) {
    // Implementation
}

// Generate docs
swag init -g cmd/server/main.go
```

### Environment Management
```bash
# Environment variables for different stages
# .env.development
DATABASE_URL=postgres://localhost:5432/chess_tournaments_dev
REDIS_URL=redis://localhost:6379
STRIPE_SECRET_KEY=sk_test_...
JWT_SECRET=dev_secret_key

# .env.production
DATABASE_URL=postgres://prod-host:5432/chess_tournaments
REDIS_URL=redis://prod-cache:6379
STRIPE_SECRET_KEY=sk_live_...
JWT_SECRET=${PRODUCTION_JWT_SECRET}
```

### Deployment Scripts
```bash
# Build and deploy script
#!/bin/bash
set -e

echo "Building backend..."
docker build -t chess-tournaments-api:latest ./backend

echo "Building frontend..."
docker build -t chess-tournaments-web:latest ./frontend

echo "Deploying to production..."
docker tag chess-tournaments-api:latest $ECR_REGISTRY/chess-tournaments-api:latest
docker push $ECR_REGISTRY/chess-tournaments-api:latest

aws ecs update-service --cluster production --service chess-tournaments-api --force-new-deployment
```

### Performance Monitoring
```go
// Custom metrics for monitoring
type Metrics struct {
    TournamentCreations prometheus.Counter
    RegistrationLatency prometheus.Histogram
    PaymentErrors       prometheus.Counter
}

func (m *Metrics) RecordTournamentCreation() {
    m.TournamentCreations.Inc()
}

func (m *Metrics) RecordRegistrationLatency(duration time.Duration) {
    m.RegistrationLatency.Observe(duration.Seconds())
}
```

This technical context provides the foundation for all development decisions and ensures consistency across the development team while maintaining scalability and performance requirements.
