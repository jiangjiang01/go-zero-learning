# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A full-stack e-commerce learning project using go-zero v1.9.3 framework with Vue 3 frontend. The project follows a single-module structure with the Go backend in `backend/` and Vue frontend in `frontend/`.

## Build & Run Commands

### Backend
```bash
# Run API server (from project root)
cd backend/service/user/api && go run user-api.go -f etc/user-api.yaml

# Generate code from .api files
cd backend/service/user/api && goctl api go -api main.api -dir ./

# Install dependencies
cd backend && go mod download
```

### Frontend
```bash
cd frontend
pnpm install          # Install dependencies
pnpm run dev          # Development server (Vite)
pnpm run build        # Production build
pnpm run lint         # ESLint
pnpm run type-check   # TypeScript checking
```

### Test Scripts
```bash
./scripts/test-order-cancel-job.sh      # Test cron job functionality
./scripts/test-permission-middleware.sh  # Test permission validation
```

## Architecture

### Backend Structure (`backend/`)
- **common/**: Shared utilities (db, jwt, middleware, errorx, response, redis, cron, ctxdata, validator)
- **model/**: GORM models (user, role, permission, menu, product, category, order, cart + associations)
- **service/user/api/**: Main API service
  - `desc/*.api`: API definition files (goctl format)
  - `internal/handler/`: HTTP handlers (auto-generated)
  - `internal/logic/`: Business logic (manually implemented)
  - `internal/svc/`: ServiceContext for dependency injection

### Key Patterns

**ServiceContext (DI)**:
```go
type ServiceContext struct {
    Config config.Config
    DB     *gorm.DB
    JWT    *jwt.JWTManager
    Redis  *redis.Redis
    Cron   *cron.CronManager
}
```

**Error Handling**: Centralized error codes (1000-9999) defined in `common/errorx/`. Use `errorx.HandleError()` in handlers and return `errorx.Err*` constants from logic.

**Middleware Stack**: LoggingMiddleware → AuthMiddleware (JWT) → PermissionMiddleware (RBAC)

**Caching**: Redis Cache-Aside pattern with random expiration (300-360s), MD5 hashed keys

### Frontend Structure (`frontend/src/`)
- **api/**: Axios API clients
- **views/system/**: Management pages (user, role, permission, menu, product, order, category, cart)
- **router/**: Route configuration

## Configuration

- **Backend config**: `backend/service/user/api/etc/user-api.yaml`
- **Database**: MySQL on `127.0.0.1:3307`, database `testdb`
- **Redis**: `127.0.0.1:6379`
- **API Port**: 8888
- **Frontend**: `http://localhost:5173` (Vite dev server)

## API Design

RESTful conventions with plural resource names:
- `POST /api/users` (create), `GET /api/users` (list), `GET /api/users/:id` (detail)
- `PUT /api/users/:id` (update), `DELETE /api/users/:id` (delete)
- Special: `GET /api/users/me` (current user), `POST /api/users/login`

Response format:
```json
{"code": 0, "message": "success", "data": {...}, "timestamp": 1705939200}
```

## Error Code Ranges

- 1000-1999: Generic errors
- 2000-2999: User errors
- 3000-3999: Role errors
- 4000-4999: Permission errors
- 5000-5999: Menu errors
- 6000-6999: Product errors
- 7000-7999: Order errors
- 8000-8999: Category errors
- 9000-9999: Cart errors

## Development Notes

- go.mod is at `backend/` (module: go-zero-learning)
- Use `127.0.0.1` instead of `localhost` in configs
- Price stored as int64 (cents/分), convert to yuan in frontend
- Test user: `admin` / `123456`
