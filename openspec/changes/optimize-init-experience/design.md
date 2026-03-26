## Context

The current init experience spans the Vue page at `web/src/view/init/index.vue`, request helpers in `web/src/api/initdb.js`, and the backend init API/service under `server/api/v1/system` and `server/service/system`. The frontend already presents database, Redis, and admin-password steps, but the backend `InitDB` request model only persists database settings. This leaves a mismatch where Redis can be tested independently but is not actually applied by the main initialization flow. The page also exposes `oracle` as a selectable type even though init routing and service selection only cover `mysql`, `pgsql`, `sqlite`, and `mssql`.

## Goals / Non-Goals

**Goals:**
- Make the init wizard reflect the actual backend-supported initialization paths.
- Ensure database and Redis "tested" state is invalidated whenever relevant inputs change.
- Extend the init contract so a successful initialization can persist optional Redis configuration together with database settings.
- Reject unsupported database types explicitly instead of silently falling back to MySQL behavior.

**Non-Goals:**
- Rework the overall init route structure or add a new multi-page setup flow.
- Introduce Oracle initialization support.
- Change the existing table/data bootstrap pipeline beyond what is needed for the init contract and config persistence.

## Decisions

### Keep the existing init endpoints and strengthen their contract

The change will continue using `/init/testDb`, `/init/testRedis`, and `/init/initdb` rather than introducing new API surface area. The backend request models will be updated so `InitDB` can carry an optional Redis configuration block and can validate supported `dbType` values up front.

Alternatives considered:
- Add a dedicated `/init/config` endpoint for Redis persistence. Rejected because it would split a single user action across two server-side workflows and complicate rollback.
- Leave `InitDB` unchanged and treat Redis as client-only setup guidance. Rejected because the UI already implies Redis participates in initialization.

### Centralize cross-cutting config writeback in the init service context

Database-specific handlers already own database connection creation and per-engine config materialization. Redis and `system.use-redis` are cross-cutting settings, so the shared init service will carry Redis config through context and persist it when initialization completes. This keeps database handlers focused on engine-specific work while allowing the final config write to remain consistent.

Alternatives considered:
- Teach every database handler to write Redis config independently. Rejected because it duplicates the same Redis/system persistence logic four times.
- Persist Redis config before tables/data initialization. Rejected because a partial failure would leave config claiming Redis is enabled even when bootstrap did not finish.

### Make frontend step state derived from current input, not past success toasts

The init page will treat successful DB/Redis tests as snapshots of the current input. When the user edits a field that affects a prior test, the corresponding success state will be cleared and the next-step/submit guards will require a fresh test. Database-type defaults will be limited to supported backends and should not silently inject insecure admin password values.

Alternatives considered:
- Keep the existing boolean flags and rely on user discipline to re-test. Rejected because it permits stale success states.
- Remove test gating entirely. Rejected because it weakens the first-run safety net.

## Risks / Trade-offs

- [More backend validation may reject flows that previously "worked" by accident] -> Return explicit validation messages so users can correct input quickly.
- [Redis persistence adds one more cross-module concern to init] -> Keep Redis writeback in shared service logic and cover it with targeted tests.
- [Resetting test state on input change can feel stricter] -> Scope invalidation only to fields that materially affect the tested connection.

## Migration Plan

No schema or data migration is required. Deploy the backend and frontend changes together so the updated init page and the expanded init request payload stay compatible. Rollback is straightforward by reverting the init page and request/service changes as one unit.

## Open Questions

- Whether the init API should return structured field-level validation details or continue using the existing message-based response wrapper.
- Whether Redis enablement should also populate `redis-list` for future multi-instance support, or only the primary `redis` block for now.
