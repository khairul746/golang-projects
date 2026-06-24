# Research & Technical Decisions: Daily Task Tracker

## Backend HTTP Router

**Decision**: Use `github.com/gorilla/mux` for routing in addition to `net/http`.
**Rationale**: Requested specifically by the user. While Go 1.22 introduced enhanced routing, `gorilla/mux` remains a proven, simple, and feature-rich router that aligns with the simplicity principle by providing straightforward path variable extraction and method filtering.
**Alternatives considered**: Go 1.22 standard `net/http` ServeMux.

## Database Driver

**Decision**: Use `github.com/lib/pq` with standard `database/sql`.
**Rationale**: `lib/pq` is the most straightforward pure-Go Postgres driver. It adheres to the simplicity principle (no ORMs, pure SQL).
**Alternatives considered**: GORM (rejected: violates "Don't over-detail" and "simple tech" principles by adding unnecessary ORM magic).

## Frontend Stack

**Decision**: Vanilla HTML, CSS, and JavaScript.
**Rationale**: Explicit user request. Eliminates build steps, bundlers, and heavy node_modules. Aligns perfectly with the constitution's simplicity requirements.
**Alternatives considered**: React, Vue (rejected: too complex for a solo developer building a simple tracker).

## Rollover Mechanism

**Decision**: Carry-over logic happens seamlessly on read/load or via a dedicated simple SQL query updating dates/statuses if necessary. Since tasks are "daily", we can simply fetch all uncompleted tasks regardless of date and display them alongside today's tasks.
**Rationale**: Avoids needing a complex cron job/background worker (which violates simplicity).
