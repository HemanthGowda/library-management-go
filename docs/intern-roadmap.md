# Intern Roadmap

This project ships with **partial** functionality on purpose. Use the implemented user registration, login, and book-update slices as the pattern library. The **Author** resource has no table and no code—design it from scratch.

Product requirements for each remaining capability are in [tasks/README.md](tasks/README.md), one file per capability. This roadmap is the suggested pace. Follow the task file for what the product must do.

Suggested pace: 1–2 milestones per week. Adjust freely.

## Roles reminder

| Role | Intended capabilities |
|------|------------------------|
| `MEMBER` | Register, browse catalog, borrow/return (future) |
| `LIBRARIAN` | Manage catalog (books/authors), manage loans |
| `ADMIN` | Everything librarians can do, plus user administration |

---

## Phase 0 — Orientation (Day 1)

**Goal:** Run the stack and understand the reference slices.

Acceptance criteria:

- [ ] `docker compose up --build` succeeds
- [ ] Swagger UI loads and login works with the seeded admin
- [ ] You can identify handler → service → repository for users and books
- [ ] You can explain how JWT middleware and role checks protect `PUT /api/v1/books/{id}`

---

## Phase 1 — Author domain design (Week 1)

**Goal:** Propose the Author resource with **no production code yet**.

Deliverables:

- [ ] Short design note (`docs/design/author.md`) covering:
  - Columns (for example name, biography, birth year, nationality)
  - Primary key, indexes, and uniqueness rules
  - Relationship to `books` (free-text `author_name` today versus `author_id` foreign key)
  - API sketch (paths, methods, roles)
- [ ] Review and approval from a mentor before coding

Acceptance criteria:

- [ ] Design answers create, read, update, delete, and list/search
- [ ] Migration plan described (up and down SQL, plus optional seed authors)
- [ ] Authorization matrix defined per endpoint

---

## Phase 2 — Author persistence + migration (Week 1–2)

**Goal:** Create the PostgreSQL table the same way `users` and `books` exist.

Deliverables:

- [ ] `000002_authors.up.sql` and `.down.sql`
- [ ] `Author` model and `Repository`
- [ ] Unit tests for any normalization helpers

Acceptance criteria:

- [ ] App starts with the new migration applied
- [ ] Unique constraints behave as designed
- [ ] No Author HTTP API yet

---

## Phase 3 — Author API vertical slice (Week 2)

**Goal:** Implement at least one full Author operation end-to-end, then complete CRUD.

Recommended order (mirror the starter style):

1. `POST /api/v1/authors`
2. `GET /api/v1/authors/{id}`
3. `GET /api/v1/authors`
4. `PUT /api/v1/authors/{id}`
5. `DELETE /api/v1/authors/{id}`

Acceptance criteria for each endpoint:

- [ ] Validation and `domain.Error` responses
- [ ] Correct role checks
- [ ] Service unit tests
- [ ] Integration tests (auth success and failure, plus validation)
- [ ] `openapi/openapi.yaml` updated

---

## Phase 4 — Wire Authors into Books (Week 3)

**Goal:** Replace or complement the free-text `author_name` on books.

Options (pick one with a mentor):

- Store `author_id` on `books` and resolve author details in responses
- Keep `author_name` denormalized and sync it when an author is renamed

Acceptance criteria:

- [ ] Migration strategy for existing seeded books
- [ ] Book update validation rejects unknown authors if you use a foreign key
- [ ] Tests cover the relationship rules

Also complete the remaining Book APIs using the update slice as a template:

- [ ] `POST /api/v1/books`
- [ ] `GET /api/v1/books/{id}`
- [ ] `GET /api/v1/books`
- [ ] `DELETE /api/v1/books/{id}`

---

## Phase 5 — Loans / borrow and return (Week 3–4)

**Goal:** Introduce a new aggregate from scratch (similar ambition to Authors).

Suggested table: `loans` with `book_id`, `member_id`, `borrowed_at`, `due_at`, `returned_at`, and `status`.

Acceptance criteria:

- [ ] A member can borrow an available copy
- [ ] Available copy count decrements and increments correctly
- [ ] Borrowing fails when no copies remain
- [ ] Return marks the loan complete and restores availability
- [ ] Librarians can list overdue loans
- [ ] Full test and OpenAPI coverage

---

## Phase 6 — User administration and hardening (Week 4+)

Stretch goals:

- [ ] Admin endpoints to list users, change roles, and enable or disable accounts
- [ ] Refresh tokens or shorter-lived access tokens
- [ ] Request logging and audit fields
- [ ] CI workflow running `go test ./...`
- [ ] Login throttling

Remember: today's JWT middleware trusts the token until expiry. A disable-user endpoint should define whether existing tokens keep working.

---

## Definition of Done (every feature)

1. Code follows existing package and naming patterns
2. Validation and authorization are explicit
3. A reversible SQL migration is added when the schema changes
4. Unit and integration tests pass locally
5. OpenAPI documents the new surface area
6. Roadmap checkboxes in this file are updated

## Reference endpoints (already done)

Use these as the gold standard:

- `POST /api/v1/users`
- `POST /api/v1/auth/login`
- `PUT /api/v1/books/{id}`
