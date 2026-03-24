# Go Boilerplate Test 🚀

A production-ready Go backend boilerplate with comprehensive features including PostgreSQL, Redis, authentication, email notifications, and observability built-in. Perfect for building scalable microservices and REST APIs.

## ✨ Features

- **🔐 Authentication** - Clerk SDK integration for secure user authentication and management
- **💾 Database** - PostgreSQL with pgx driver and automatic migrations using Tern
- **⚡ Caching & Job Queue** - Redis integration with Asynq for async task processing
- **📧 Email Service** - Built-in email notifications with Resend integration
- **📊 Observability** - New Relic APM integration with distributed tracing and performance monitoring
- **📝 Structured Logging** - Zerolog for structured JSON logging with multiple levels
- **🛡️ Comprehensive Middleware** - Request ID tracking, rate limiting, context management, tracing
- **❤️ Health Checks** - Database and Redis readiness endpoints for orchestration
- **🔄 Error Handling** - Custom error types and comprehensive error management
- **⚙️ Configuration Management** - Environment-based configuration with validation
- **🧪 Testing Utilities** - Testing helpers with testcontainers support
- **📚 OpenAPI/Swagger** - Auto-generated API documentation
- **🐳 Docker Support** - Docker Compose for seamless local development

## 🏗️ Architecture Overview

```
┌─────────────────────────────────────────┐
│         Echo Web Framework              │
├─────────────────────────────────────────┤
│  Middleware Layer (Auth, Tracing, etc)  │
├─────────────────────────────────────────┤
│  Handler Layer (HTTP Endpoints)         │
├─────────────────────────────────────────┤
│  Service Layer (Business Logic)         │
├─────────────────────────────────────────┤
│  Repository Layer (Data Access)         │
├─────────────────────────────────────────┤
│  PostgreSQL + Redis + External APIs     │
└─────────────────────────────────────────┘
```

## 🏗️ Tech Stack

### Backend (Go)
| Component | Technology | Version |
|-----------|-----------|---------|
| Web Framework | Echo | v4 |
| Database | PostgreSQL | 16 Alpine |
| Cache/Queue | Redis | Latest |
| Database Driver | pgx | v5 |
| Async Jobs | Asynq | v0.26.0 |
| Authentication | Clerk SDK | v2 |
| Email | Resend | v2 |
| Logging | Zerolog | v1.34.0 |
| Validation | go-playground/validator | v10 |
| Configuration | Koanf | v2 |
| Migrations | Tern | v2 |
| Observability | New Relic Agent | v3 |
| Testing | Testify, Testcontainers | Latest |

### Frontend Packages (TypeScript/Node)
- Email templates (React/TSX)
- OpenAPI contract generation
- Shared Zod validation schemas

## 📋 Prerequisites

- **Go:** 1.25.0 or higher
- **Docker:** Latest with Docker Compose support
- **Task:** For running development tasks (or Make)
- **Node.js/Bun:** For frontend packages (optional)

## 🚀 Quick Start

### 1. Clone & Setup

```bash
git clone https://github.com/ShiunduZachariah/go-boilerplate-test.git
cd go-boilerplate-test
```

### 2. Start Services

```bash
# Start PostgreSQL (port 5434) and Redis (port 6379)
docker-compose up -d

# Verify services are running
docker ps
```

### 3. Configure Environment

```bash
cd backend

# Copy environment template (if needed)
cp .env.sample .env

# Edit .env with your configuration:
# - Database credentials
# - Clerk authentication keys
# - Resend API key
# - New Relic license key
# - Other service credentials
```

### 4. Run Application

```bash
# Start the Go server
task run

# Or directly:
go run ./cmd/go-boilerplate-test
```

Server starts at `http://localhost:8080`

## 📁 Project Structure

```
go-boilerplate-test/
├── backend/
│   ├── cmd/
│   │   └── go-boilerplate-test/          # Application entry point
│   ├── internal/                         # Private packages
│   │   ├── config/                       # Configuration & environment management
│   │   ├── database/                     # DB connections & migrations
│   │   ├── handler/                      # HTTP request handlers
│   │   ├── middleware/                   # Custom middleware
│   │   ├── router/                       # Route definitions
│   │   ├── server/                       # Server initialization
│   │   ├── service/                      # Business logic layer
│   │   ├── repository/                   # Data access layer
│   │   ├── logger/                       # Logging configuration
│   │   ├── lib/
│   │   │   ├── email/                    # Email service
│   │   │   ├── job/                      # Background jobs (Asynq)
│   │   │   └── utils/                    # Utility functions
│   │   ├── testing/                      # Testing utilities
│   │   ├── errs/                         # Custom error types
│   │   ├── sqlerr/                       # SQL error handling
│   │   └── model/                        # Data models
│   ├── static/
│   │   └── openapi.json                  # API specification
│   ├── go.mod & go.sum                   # Go dependencies
│   ├── .env                              # Environment variables
│   ├── .env.sample                       # Configuration template
│   ├── Taskfile.yml                      # Development tasks
│   ├── golangci.yml                      # Linter configuration
│   └── README.md                         # Backend documentation
│
├── packages/
│   ├── emails/                           # Email templates (React/TSX)
│   ├── openapi/                          # API contract generation
│   └── zod/                              # Shared validation schemas
│
├── docker-compose.yml                    # Local development services
├── turbo.json                            # Monorepo configuration
├── package.json                          # Root package config
└── README.md                             # This file
```

## ⚙️ Configuration

### Environment Variables Structure

Configuration uses the `BOILERPLATE_` prefix:

```dotenv
# Environment
BOILERPLATE_PRIMARY.ENV="local|development|production"

# Server Configuration
BOILERPLATE_SERVER.PORT="8080"
BOILERPLATE_SERVER.READ_TIMEOUT="30"
BOILERPLATE_SERVER.WRITE_TIMEOUT="30"
BOILERPLATE_SERVER.IDLE_TIMEOUT="60"
BOILERPLATE_SERVER.CORS_ALLOWED_ORIGINS="https://yourdomain.com"

# Database Configuration
BOILERPLATE_DATABASE.HOST="localhost"
BOILERPLATE_DATABASE.PORT="5434"
BOILERPLATE_DATABASE.USER="your_db_user"
BOILERPLATE_DATABASE.PASSWORD="your_secure_password"
BOILERPLATE_DATABASE.NAME="your_database_name"
BOILERPLATE_DATABASE.SSL_MODE="disable|require"
BOILERPLATE_DATABASE.MAX_OPEN_CONNS="25"
BOILERPLATE_DATABASE.MAX_IDLE_CONNS="5"
BOILERPLATE_DATABASE.CONN_MAX_LIFETIME="300"
BOILERPLATE_DATABASE.CONN_MAX_IDLE_TIME="300"

# Redis Configuration
BOILERPLATE_REDIS.ADDRESS="localhost:6379"

# Authentication
BOILERPLATE_AUTH.SECRET_KEY="your_secret_key"

# Integrations
BOILERPLATE_INTEGRATION.RESEND_API_KEY="your_resend_key"

# Observability
BOILERPLATE_OBSERVABILITY.LOGGING.LEVEL="debug|info|warn|error"
BOILERPLATE_OBSERVABILITY.LOGGING.FORMAT="console|json"
BOILERPLATE_OBSERVABILITY.LOGGING.SLOW_QUERY_THRESHOLD="100ms"
BOILERPLATE_OBSERVABILITY.NEW_RELIC.LICENSE_KEY="your_license_key"
BOILERPLATE_OBSERVABILITY.NEW_RELIC.APP_LOG_FORWARDING_ENABLED="true|false"
BOILERPLATE_OBSERVABILITY.NEW_RELIC.DISTRIBUTED_TRACING_ENABLED="true|false"
```

See `.env.sample` for a complete template.

## 🔌 API Endpoints

### Core Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Basic health check |
| `GET` | `/openapi` | OpenAPI/Swagger UI |
| `GET` | `/openapi.json` | OpenAPI specification |

### Health Checks
| Endpoint | Purpose |
|----------|---------|
| `GET /health/database` | PostgreSQL connectivity |
| `GET /health/redis` | Redis connectivity |
| `GET /health/full` | Complete system status |

## 🧪 Testing

```bash
cd backend

# Run all tests
go test ./...

# Run tests with coverage
go test ./... -cover

# Run specific package tests
go test ./internal/handler -v

# Run with verbose output
go test ./... -v -count=1

# Run benchmarks
go test -bench=. -benchmem ./...
```

## 🗄️ Database

### Automatic Migrations

Migrations run automatically on startup. Create new migrations:

1. Add SQL file to `backend/internal/database/migrations/`
2. Name format: `NNN_description.sql` (e.g., `002_create_users_table.sql`)
3. Follow the split marker format:

```sql
-- migrations/002_create_users_table.sql

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

---- create above / drop below ----

DROP TABLE IF EXISTS users;
```

### Database Connection Details

| Setting | Value |
|---------|-------|
| **Host** | localhost |
| **Port** | 5434 |
| **User** | Set in `.env` |
| **Password** | Set in `.env` |
| **Database** | Set in `.env` |
| **SSL Mode** | disable (local) / require (production) |

## 📧 Email Service

Send emails through Resend integration:

```go
type EmailService interface {
    SendWelcomeEmail(ctx context.Context, email, name string) error
    SendPasswordReset(ctx context.Context, email, resetLink string) error
    SendNotification(ctx context.Context, email, subject, body string) error
}
```

Templates are React components in `packages/emails/src/templates/`

## ⚡ Background Jobs

Process jobs asynchronously with Asynq:

```go
// Enqueue tasks
taskService := services.JobService
err := taskService.EnqueueEmailTask(ctx, "user@example.com", "welcome")

// Jobs process in background via Redis
```

Handlers: `internal/lib/job/handlers.go`

## 🔐 Authentication

Clerk integration for user authentication:

```go
// Access authenticated user from context
userID := ctx.Get("user_id")
userEmail := ctx.Get("user_email")
```

Protected routes use the auth middleware automatically.

## 📊 Observability

### Logging

Structured JSON logging with Zerolog:

```go
import "github.com/rs/zerolog/log"

log.Info().Msg("Server started")
log.Error().Err(err).Msg("Database connection failed")
log.Debug().Interface("data", payload).Msg("Request received")
```

### New Relic APM

Real-time performance monitoring:
- Application metrics
- Transaction tracing
- Error tracking
- Custom events
- Distributed tracing

Configure in `.env` with your license key.

### Health Checks

Monitor system status via health endpoints:

```bash
# Quick health check
curl http://localhost:8080/health

# Database status
curl http://localhost:8080/health/database

# Redis status
curl http://localhost:8080/health/redis
```

## 🐳 Docker

### Container Management

```bash
# Start services (background)
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Remove volumes and reset data
docker-compose down -v

# Restart services
docker-compose restart
```

### Individual Container Commands

```bash
# PostgreSQL
docker logs my-postgres
docker exec my-postgres psql -U user -d database -c "SELECT 1;"

# Redis
docker logs my-redis
docker exec my-redis redis-cli ping
```

## 🔧 Development Tasks

Using Task (Taskfile.yml):

```bash
# Start development server (watch mode)
task run

# Run database migrations
task migrate

# Run test suite
task test

# Run linter
task lint

# Format code
task fmt

# Build binary
task build
```

## 📚 API Documentation

Access interactive API documentation:

```
http://localhost:8080/openapi
```

Or get raw OpenAPI spec:

```
http://localhost:8080/openapi.json
```

## 🚨 Error Handling

Comprehensive error management with custom types:

```go
// Custom error responses
// Proper HTTP status codes
// Structured error information
// Error logging and tracking
```

Error types: `internal/errs/`
SQL errors: `internal/sqlerr/`

## 🛠️ Troubleshooting

### PostgreSQL Not Connecting

```
Error: failed to connect to database
```

**Solutions:**
```bash
# Verify containers running
docker ps

# Check logs
docker logs my-postgres

# Verify .env credentials match docker-compose.yml
# Default: user=appuser, password=apppass, db=appdb, port=5434

# Restart containers
docker-compose down && docker-compose up -d
```

### Redis Connection Failed

```
Error: failed to connect to redis
```

**Solutions:**
```bash
# Verify Redis is running
docker ps | grep redis

# Test Redis connection
docker exec my-redis redis-cli ping

# Check BOILERPLATE_REDIS.ADDRESS in .env
```

### Port Already in Use

```
Error: listen tcp :8080: bind: address already in use
```

**Solutions:**
```bash
# Change port in .env
BOILERPLATE_SERVER.PORT="8081"

# Or find and kill process
lsof -i :8080
kill -9 <PID>
```

### Import/Build Errors

**Solutions:**
```bash
# Download dependencies
go mod download

# Tidy modules
go mod tidy

# Clean build cache
go clean -cache

# Rebuild
go build ./cmd/go-boilerplate-test
```

## 📝 Code Quality

### Standards & Tools

- **Linter:** golangci-lint (`.golangci.yml`)
- **Format:** gofmt / goimports
- **Testing:** Testify, Testcontainers
- **Benchmarking:** Go built-in

### Pre-commit Checklist

```bash
# Format code
task fmt

# Run linter
task lint

# Run tests
task test

# Build check
go build ./...
```

## 🤝 Contributing

1. Create feature branch: `git checkout -b feature/amazing-feature`
2. Make changes and test
3. Run linter and tests: `task lint && task test`
4. Commit: `git commit -m 'Add amazing feature'`
5. Push: `git push origin feature/amazing-feature`
6. Open Pull Request

### Code Style

- Follow Go conventions
- Use meaningful variable names
- Add comments for exported functions
- Write tests for new features
- Keep functions focused and small

## 📄 License

MIT License - see LICENSE file for details

## 👤 Author

[ShiunduZachariah](https://github.com/ShiunduZachariah)

## 🆘 Support & Feedback

- 🐛 Found a bug? [Open an issue](https://github.com/ShiunduZachariah/go-boilerplate-test/issues)
- 💡 Have a suggestion? [Discussions](https://github.com/ShiunduZachariah/go-boilerplate-test/discussions)
- 📧 Questions? Contact via GitHub

## 🙏 Acknowledgments

Built with modern Go best practices and production-ready patterns.

---

**Made with ❤️ for the Go community**

Last Updated: March 2026
