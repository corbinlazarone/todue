# AGENTS

## Environment
- Copy `.env.example` to `.env` before running anything; every Make target (`make run`, migrations, goose) and the Docker services read it, so missing/incorrect values (DB_CONN, Google/OIDC creds, `SESSION_PASSWORD`, `NEXT_PUBLIC_API_URL`) block both frontend and backend.
- `make run` loads `.env`, but it also rewrites `DB_CONN` by replacing `@db:` with `@localhost:` because the Go server expects a local host rather than the Docker host alias.

## Architecture/Services
- Local stack always pairs `web-client/` (Next.js 16 + React 19) on port 3000, the Go API at `server/cmd/api` listening on 4000, and Postgres 16 on 5432; Docker Compose wires them together with `NEXT_PUBLIC_API_URL=http://server:4000` so the front-end hits the containerized backend.
- When running with `make dev`/`dev-build`, `docker-compose.yml` sources the shared `.env`, so `NEXT_PUBLIC_API_URL`, `SESSION_PASSWORD`, and Google client IDs must exist in that file for the container builds to succeed.

## Common workflows
- `make dev` brings up all services; use `make dev-build` after code changes that touch Docker images, and `make dev-stop`/`make dev-restart` to stop or restart without rebuilding. `make dev-logs` follows logs for every service, `make dev-logs-server` tails only the API container.
- For backend-only runs, bring the DB up first (`make db-up`), then run `make run` from the repo root (it expects `.env`). The command fails fast if `.env` is missing.
- Database helpers: `make db-up`/`db-down` just start/stop Postgres; `make db-reset` prompts for confirmation before destroying volumes; `make db-shell` shells into Postgres using credentials from `.env`.

## Testing, build, and migrations
- `make test` (server-only) executes `go test ./...` inside `server`. `make build` runs the same tests and then builds `server/cmd/api` into `bin/todue-server`.
- Migration targets run `goose` from `server/cmd/internals/migrations`, pulling the `DB_CONN` string from `.env` and swapping `@db:` for `@localhost:`; `make migrate-up`/`migrate-down`/`migrate-status` rely on that rewrite and the goose binary being installed or available in the Go module context.

## Frontend dev notes
- When not using Docker, run the Next.js app with `cd web-client && npm install && npm run dev` (Node 18+). The app expects the API URL to live in `NEXT_PUBLIC_API_URL`; keep it pointed at `http://localhost:4000` when hitting the locally running Go server.
