## Why

The current init flow presents a polished multi-step UI, but its behavior does not fully match the backend contract: connection checks can become stale after form edits, unsupported database choices are exposed, and Redis settings submitted from the page are not applied during initialization. This creates avoidable friction in the first-run experience and makes the init page unreliable as the source of truth for system bootstrap.

## What Changes

- Refine the init wizard so each step reflects real backend capabilities, surfaces clearer guidance, and keeps validation/test status in sync with the current form state.
- Align database type choices and default values with the database handlers actually supported by the init backend.
- Extend the init request/response contract so the backend can consume optional Redis settings during initialization instead of ignoring them.
- Update init service logic to validate, test, and persist both database and Redis configuration consistently when initialization succeeds.
- Add or update automated coverage for init page interactions and backend init request handling.

## Capabilities

### New Capabilities
- `system-initialization-flow`: A first-run initialization flow that guides database and optional Redis setup with UI state and backend behavior kept in sync.

### Modified Capabilities

## Impact

- Frontend: `web/src/view/init/index.vue`, `web/src/api/initdb.js`
- Backend API: `server/api/v1/system/sys_initdb.go`, `server/router/system/sys_initdb.go`
- Backend request/service/config flow: `server/model/system/request/sys_init.go`, `server/service/system/sys_initdb*.go`
- System configuration writeback for database and Redis settings
- Test coverage for init submission, connection tests, and step-state behavior
