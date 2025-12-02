# Todue - Agent Development Guide

## Build & Test Commands

### Frontend (Next.js)
- `npm run dev` - Start development server
- `npm run build` - Build for production  
- `npm run lint` - Run ESLint

### Backend (Go)
- `make run` - Run server locally (requires DB)
- `make build` - Build server binary
- `go test ./...` - Run all tests
- `go test -run TestSpecific ./pkg` - Run single test

### Docker Development
- `make dev` - Start all services
- `make dev-build` - Rebuild and start services
- `make migrate-up/down` - Database migrations

## Code Style Guidelines

### Go Backend
- **Naming**: PascalCase for types, camelCase for functions/variables
- **Error Handling**: Explicit `if err != nil` checks, structured logging
- **Packages**: Organize by domain (models, auth, migrations)
- **JSON Tags**: camelCase, **DB Tags**: snake_case
- **Dependencies**: Inject via `application` struct

### Next.js Frontend  
- **Components**: PascalCase, feature-based organization
- **Imports**: External libs first, then internal with `@/` alias
- **Styling**: Tailwind CSS + shadcn/ui components, use `cn()` utility
- **Types**: Strict TypeScript, interfaces for props
- **Files**: `.tsx` for components, `.ts` for utilities

### Error Patterns
- **Backend**: Consistent JSON responses with `type` and `message` fields
- **Frontend**: Toast notifications with `sonner`, try-catch for async

### Database
- Use `pgx/v5` with connection pooling
- Context-aware operations
- Timestamped migrations with goose