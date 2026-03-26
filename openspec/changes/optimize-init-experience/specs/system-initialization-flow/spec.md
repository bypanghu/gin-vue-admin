## ADDED Requirements

### Requirement: Init wizard SHALL expose only supported database bootstrap options
The system SHALL present only database types that the initialization backend can test and initialize successfully. When the user switches database type, the form SHALL apply matching database defaults without introducing unsupported values or silently overwriting unrelated setup fields.

#### Scenario: Supported database options are shown
- **WHEN** a user opens the init wizard database step
- **THEN** the database type selector lists only backends supported by the init service

#### Scenario: Switching database type updates only relevant defaults
- **WHEN** a user changes the selected database type
- **THEN** the database connection fields update to defaults for that backend and previously entered Redis/admin settings remain intact

### Requirement: Connection test state SHALL reflect the current input snapshot
The system SHALL treat successful database and Redis connection tests as valid only for the exact input values that were tested. Any change to a field that affects a tested connection MUST clear the corresponding success state until the user tests again.

#### Scenario: Database edits invalidate prior database test success
- **WHEN** a user changes a database connection field after a successful database test
- **THEN** the init wizard marks the database test as no longer valid and blocks progression that requires a passing database test

#### Scenario: Redis edits invalidate prior Redis test success
- **WHEN** a user changes an enabled Redis connection field after a successful Redis test
- **THEN** the init wizard marks the Redis test as no longer valid and requires a fresh Redis test before final submission

### Requirement: Initialization SHALL apply optional Redis configuration together with database setup
The system SHALL accept an optional Redis configuration block as part of the init submission. When Redis is enabled and the configuration is valid, the backend SHALL persist the Redis settings and enable Redis in system configuration as part of the successful initialization flow.

#### Scenario: Redis-enabled initialization persists Redis settings
- **WHEN** a user submits init with a valid database configuration and Redis enabled
- **THEN** the backend writes both the selected database configuration and the Redis enablement/settings to the generated system config

#### Scenario: Redis-disabled initialization leaves Redis disabled
- **WHEN** a user submits init with Redis disabled
- **THEN** the backend completes initialization without requiring Redis fields and persists Redis as disabled in system configuration

### Requirement: Init APIs SHALL reject unsupported database types explicitly
The system MUST reject unsupported database types for database testing and initialization with a clear error response instead of silently routing them to another database handler.

#### Scenario: Unsupported database type is rejected during test
- **WHEN** a client calls the database test API with an unsupported `dbType`
- **THEN** the API returns a failure response explaining that the database type is not supported for initialization

#### Scenario: Unsupported database type is rejected during init
- **WHEN** a client calls the init API with an unsupported `dbType`
- **THEN** the API returns a failure response and does not attempt database creation or config writeback
