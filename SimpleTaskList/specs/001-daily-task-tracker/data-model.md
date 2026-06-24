# Data Model: Daily Task Tracker

## Entities

### `Task`

Represents a single unit of work in the daily task list.

**Fields**:
- `id` (UUID or Serial Integer): Primary Key.
- `title` (String): The name of the task. Max length 255.
- `status` (String/Enum): The current state. Allowed values: `NOT_STARTED`, `IN_PROGRESS`, `DONE`. Default: `NOT_STARTED`.
- `estimated_duration` (String/Integer): The estimated time (e.g., "30m", "2h", or stored as minutes in integer). Let's use `integer` representing minutes for simplicity.
- `created_at` (Timestamp): Record creation time.
- `updated_at` (Timestamp): Record update time.

## Relationships

None. Single-table design to adhere to the simplistic constitution.

## Validation Rules

- `title` must not be empty.
- `status` must be one of the three allowed states.
- `estimated_duration` cannot be negative.

## State Transitions

- `NOT_STARTED` -> `IN_PROGRESS`
- `IN_PROGRESS` -> `DONE`
- `NOT_STARTED` -> `DONE` (allowed for quick completion)
- Any state can revert to `NOT_STARTED` if a mistake was made.
