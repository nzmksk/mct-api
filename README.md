# MY Chess Tour - API

A Go-based REST API for managing chess tournaments in Malaysia and the ASEAN region.

## Features

- **Tournament Management**: Create, manage, and execute chess tournaments
- **User Authentication**: JWT-based authentication for players and organizers
- **Payment Processing**: CHIP integration with escrow system
- **Real-time Updates**: WebSocket support for live tournament updates
- **Mobile Optimized**: API designed for mobile-first applications

## Tech Stack

- **Backend**: Go 1.21+ with Gin web framework
- **Database**: PostgreSQL 15+ with Redis for caching
- **Authentication**: JWT tokens with bcrypt password hashing
- **Payments**: CHIP Collect for entry fee payments and CHIP Send for organizers' payouts
- **Deployment**: Docker containers with AWS ECS

## Getting Started

### Prerequisites

- Go 1.21 or later
- PostgreSQL 15+
- Redis 7+
- Docker (optional)

### Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd mct-api
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Start databases with Docker**
   ```bash
   docker-compose up postgres redis -d
   ```

5. **Run database migrations**
   ```bash
   # Migration commands will be added when migrate tool is set up
   ```

6. **Start the server**
   ```bash
   go run cmd/server/main.go
   ```

The API will be available at `http://localhost:8080`

### Using Docker

1. **Start all services**
   ```bash
   docker-compose --profile full-stack up -d
   ```

2. **View logs**
   ```bash
   docker-compose logs -f api
   ```

3. **Stop services**
   ```bash
   docker-compose down
   ```

## API Endpoints

### Health Check
- `GET /health` - Service health status
- `GET /api/v1/ping` - API connectivity test

### Authentication (Planned)
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Token refresh
- `POST /api/v1/auth/logout` - User logout

### Tournaments (Planned)
- `GET /api/v1/tournaments` - List tournaments
- `POST /api/v1/tournaments` - Create tournament
- `GET /api/v1/tournaments/{id}` - Get tournament details
- `PUT /api/v1/tournaments/{id}` - Update tournament
- `DELETE /api/v1/tournaments/{id}` - Cancel tournament

## Project Structure

```
mct-api/
├── cmd/
│   └── server/          # Application entrypoint
├── internal/            # Private application code
│   ├── auth/           # Authentication service
│   ├── users/          # User management
│   ├── tournaments/    # Tournament management
│   ├── registrations/  # Player registration
│   ├── payments/       # Payment processing
│   ├── notifications/  # Email/SMS notifications
│   ├── analytics/      # Usage analytics
│   └── shared/         # Shared utilities
├── pkg/                # Public packages
│   ├── database/       # Database connections
│   ├── middleware/     # HTTP middleware
│   └── utils/          # Utility functions
├── migrations/         # Database migrations
├── configs/           # Configuration files
└── docs/              # Documentation
```

## Environment Variables

See `.env.example` for all available configuration options.

Key variables:
- `ENV`: Environment mode (development/production)
- `PORT`: Server port (default: 8080)
- `DB_*`: PostgreSQL connection settings
- `REDIS_*`: Redis connection settings
- `JWT_SECRET`: Secret key for JWT tokens
- `CHIP_SECRET_KEY`: CHIP API key for payments

## Development

### Code Standards

- Use `gofmt` for code formatting
- Follow Go naming conventions
- Write unit tests for all business logic
- Use structured logging with logrus

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detection
go test -race ./...
```

### Database Migrations

```bash
# Create new migration
migrate create -ext sql -dir migrations add_new_table

# Run migrations
migrate -path migrations -database $DATABASE_URL up

# Rollback migration
migrate -path migrations -database $DATABASE_URL down 1
```

## Deployment

The application is containerized and ready for deployment to AWS ECS or any Docker-compatible platform.

### Production Environment

1. Set up PostgreSQL RDS instance
2. Set up ElastiCache Redis cluster
3. Configure environment variables
4. Deploy using Docker container

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is proprietary software for MY Chess Tour platform.

## Support

For development questions or issues, please refer to the project documentation or contact the development team.
