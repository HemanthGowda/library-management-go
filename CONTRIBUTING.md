# Contributing

This repository is a teaching codebase. Prefer clarity and consistency with the existing slices over clever abstractions.

## Before you start

1. Read [README.md](README.md) for architecture and run instructions.
2. Read [docs/intern-roadmap.md](docs/intern-roadmap.md) and pick the next incomplete milestone.
3. Skim the reference implementations:
   - User registration: `internal/user`
   - Book update: `internal/book`
   - Auth/JWT: `internal/auth`
   - Errors: `internal/domain` and `internal/httpx`
   - Schema changes: `internal/db/migrations`

## Branching

- Create a feature branch from `main`: `feature/<short-description>`
- Keep pull requests focused on one milestone when possible

## Implementation checklist

For every new API:

- [ ] SQL migration (up and down) when the schema changes
- [ ] Domain model and repository methods
- [ ] Request/response structs with validation
- [ ] Service with explicit business rules
- [ ] Handler registered in `internal/httpserver`
- [ ] Role checks against the JWT principal when needed
- [ ] Service unit tests
- [ ] Integration test covering the happy path plus auth and validation failures
- [ ] OpenAPI and roadmap updates

## Coding conventions

- Go 1.26, constructor injection, [Fiber](https://github.com/gofiber/fiber) v3 routes
- Package by feature (`user`, `book`, and later `author`, `loan`)
- Do not put business rules in handlers or SQL repositories
- Never commit secrets; use `.env` or environment variables
- Do not log passwords or raw JWTs

## Tests

```bash
go test ./...
```

Docker must be available for Testcontainers-based integration tests.

## Documentation

- Update `openapi/openapi.yaml`
- Keep error responses aligned with `domain.Error`
- Mark completed roadmap items in `docs/intern-roadmap.md`

## Review expectations

Reviewers will look for:

1. Pattern parity with the existing slices
2. Validation and authorization coverage
3. Tests that would catch regressions
4. A reversible SQL migration when persistence changes
