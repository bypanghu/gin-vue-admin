## 1. Backend Init Contract

- [x] 1.1 Extend the init request models to carry optional Redis configuration and validate supported `dbType` values explicitly.
- [x] 1.2 Update the init API/service flow so successful initialization persists database config plus Redis enablement/settings in one writeback path.
- [x] 1.3 Keep `/init/testDb` and `/init/initdb` from silently falling back to MySQL when an unsupported database type is submitted.

## 2. Frontend Init Experience

- [x] 2.1 Refine `web/src/view/init/index.vue` so the database selector only shows backend-supported options and uses matching defaults.
- [x] 2.2 Reset database and Redis test-success state whenever the user edits fields that affect the tested connection.
- [x] 2.3 Update init submission payload shaping and step messaging so the UI reflects the backend-backed initialization contract.

## 3. Verification

- [x] 3.1 Add or update backend tests for supported/unsupported database type handling and Redis-aware init config persistence.
- [x] 3.2 Add or update frontend coverage for step gating, test-state invalidation, and final init submission payload behavior.
- [x] 3.3 Run the relevant frontend and backend test commands for the init flow changes.
