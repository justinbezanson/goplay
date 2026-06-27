# Phase 0: Architecture & Planning Document

## Project Summary

Turn the existing Go/Gin/Templ sample into a reusable boilerplate for new web apps.

## Tech Stack

| Layer | Choice |
|---|---|
| Language | Go 1.22+ |
| HTTP framework | Gin |
| Template engine | Templ (SSR only) |
| Database | SQLite |
| Query layer | sqlx |
| CSS | Tailwind CSS v4 (standalone CLI) |
| Auth | Session-based (encrypted cookies, sessions in SQLite) |
| Config | godotenv (.env files) |
| Migrations | golang-migrate (SQL files) |
| Email (dev) | MailHog (via docker-compose) |
| Testing | Go standard testing |
| Task runner | Makefile |
| Container | Dockerfile + docker-compose |
| CI | GitHub Actions |

## Features (included)

1. **Multi-tenant user management** — organizations with primary users who can manage members
2. **Error handling middleware** — centralized error handling with user-friendly error pages
3. **Request logging** — structured logging via `log/slog`
4. **Config management** — .env file via godotenv
5. **Makefile** — common commands (dev, build, test, migrate, tailwind)
6. **Docker** — multi-stage Dockerfile + docker-compose with MailHog for dev
7. **CI** — GitHub Actions: lint, test, build

## Features (explicitly excluded)

- Health check endpoint

## Project Structure

```
goplay/
├── main.go                 # Entry point, router setup
├── .env.example            # Env var template
├── .gitignore
├── Makefile
├── Dockerfile
├── docker-compose.yml
├── go.mod / go.sum
├── .github/workflows/
│   └── ci.yml
├── migrations/
│   ├── 000001_create_organizations.up.sql
│   ├── 000001_create_organizations.down.sql
│   ├── 000002_create_users.up.sql
│   └── 000002_create_users.down.sql
├── handlers/
│   ├── auth.go             # Register, login, logout, password reset
│   ├── home.go             # Home page
│   └── org.go              # Org settings, manage users, invite users
├── middleware/
│   ├── auth.go             # Session loading, require auth
│   └── error.go            # Panic recovery + friendly error pages
├── models/
│   ├── db.go               # DB init, migration runner
│   ├── org.go              # Organization struct + queries
│   └── user.go             # User struct + queries
├── views/
│   ├── layout.templ        # Base layout
│   ├── home.templ          # Home/welcome page
│   ├── org/
│   │   ├── settings.templ  # Org profile/settings
│   │   ├── members.templ   # User list for org
│   │   ├── invite.templ    # Invite new user form
│   │   └── direct_create.templ # Direct create user form
│   ├── auth/
│   │   ├── login.templ     # Login form
│   │   ├── register.templ  # Registration form (creates org + primary user)
│   │   └── reset.templ     # Password reset flow
│   └── errors.templ        # Error pages (404, 500)
├── static/
│   └── css/
│       └── style.css       # Tailwind output (generated)
└── planning/
    └── design.md           # This file
```

## Key Design Decisions

### Data Model

**Organization** — created during registration alongside the primary user.
| Field | Type |
|---|---|
| id | INTEGER PRIMARY KEY |
| name | TEXT |
| email | TEXT |
| address | TEXT |
| address2 | TEXT (nullable) |
| city | TEXT |
| state | TEXT |
| zip | TEXT |
| country | TEXT |
| created_at | DATETIME |

**User** — each user belongs to one organization.
| Field | Type |
|---|---|
| id | INTEGER PRIMARY KEY |
| org_id | INTEGER FK → organizations |
| email | TEXT UNIQUE |
| password_hash | TEXT |
| name | TEXT |
| role | TEXT — `"primary"` or `"user"` |
| created_at | DATETIME |

**Invitation** — tracks pending invites sent by the primary user.
| Field | Type |
|---|---|
| id | INTEGER PRIMARY KEY |
| org_id | INTEGER FK |
| email | TEXT |
| token | TEXT UNIQUE |
| expires_at | DATETIME |
| used | BOOLEAN |
| created_by | INTEGER FK → users |

### Authentication Flow
- **Registration**: creates Organization + primary User in a transaction. Primary user auto-logged in.
- **Primary user** (role = `"primary"`): the user who registered the org. Can view org settings, invite new users, directly create users, reset member passwords.
- **Regular users** (role = `"user"`): cannot manage org or other users.
- **Invite flow**: primary user enters email → system generates invitation token → email sent with accept link → recipient creates their account tied to the org.
- **Direct create**: primary user fills name + email + temporary password → user is created immediately (no invite email needed).
- **Password reset**: self-service via email (all users), plus primary user can manually reset any member's password.
- Sessions stored in SQLite. bcrypt for password hashing. MailHog for dev email.

### Database & Migrations
- SQLite file stored at `data/app.db` (gitignored)
- Migrations via `golang-migrate` with `migrations/` SQL files
- Auto-run migrations on startup

### Tailwind Integration
- Standalone Tailwind CLI watches `.templ` and `.go` files for class names
- Outputs to `static/css/style.css`
- Makefile target wraps the CLI command

### Docker
- **Dockerfile**: multi-stage (build stage with Go + templ + tailwind, runtime stage scratch/alpine)
- **docker-compose**: app + MailHog, mounts source for hot-reload

### Error Handling
- Custom recovery middleware replaces Gin's default
- Renders user-friendly error pages via Templ (404, 500)
- Logs errors via slog with stack traces

## Phases Ahead

| Phase | Description |
|---|---|
| 1 | Data Modeling — User struct, DB schema, session schema |
| 2 | Interface & API Contracts — Route declarations, handler stubs |
| 3 | TODO Mapping — Insert TODOs at every implementation point |
| 4 | Dry Run — Tentative implementation, test, revert |
| 5 | Invariants — Validation rules, security constraints |
| 6 | Final Execution — Full implementation |
