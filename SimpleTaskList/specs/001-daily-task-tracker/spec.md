# Feature Specification: Daily Task Tracker

**Feature Branch**: `001-daily-task-tracker`

**Created**: 2026-06-10

**Status**: Draft

**Input**: User description: "saya ingin membangun sebuah aplikasi simple task list yang bisa memanage task sehari-hari dengan gampang. Setiap task di track mana yang belum selesai, sedang dikerjakan, dan sudah selesai, serta alokasi waktu pada masing-masing task"

## Clarifications

### Session 2026-06-10

- Q: Format Alokasi Waktu (Time Allocation Format) → A: Estimasi Durasi (misal: 2 jam atau 30 menit)
- Q: Perilaku Pergantian Hari (Daily Rollover Behavior) → A: Otomatis terbawa ke hari berikutnya

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Create and View Daily Tasks (Priority: P1)

As a user, I want to create tasks for my day and view them in a simple list so that I know what I need to do.

**Why this priority**: Core functionality; without creating and viewing tasks, the application cannot function.

**Independent Test**: Can be fully tested by creating multiple tasks and verifying they appear correctly in the default "Not Started" state.

**Acceptance Scenarios**:

1. **Given** I am on the task list view, **When** I add a new task with a name and time allocation, **Then** the task appears in the "Not Started" list.
2. **Given** I have existing tasks, **When** I open the application, **Then** I see all my tasks organized by their status.

---

### User Story 2 - Update Task Status (Priority: P1)

As a user, I want to change the status of a task from "Not Started" to "In Progress" or "Done" so that I can track my progress throughout the day.

**Why this priority**: Essential for the core value proposition of tracking task progress.

**Independent Test**: Can be tested by creating a task and sequentially moving it through the available statuses.

**Acceptance Scenarios**:

1. **Given** a task is "Not Started", **When** I mark it as started, **Then** it moves to the "In Progress" category.
2. **Given** a task is "In Progress", **When** I mark it as completed, **Then** it moves to the "Done" category.

---

### User Story 3 - Manage Time Allocation (Priority: P2)

As a user, I want to set and view the allocated time for each task so that I can manage my daily schedule effectively.

**Why this priority**: While important, the basic task tracking can work without time allocation, making this secondary.

**Independent Test**: Can be tested by verifying time inputs are saved correctly and displayed alongside the task.

**Acceptance Scenarios**:

1. **Given** I am creating or editing a task, **When** I input a specific time allocation, **Then** the task saves and displays that time allocation.

### Edge Cases

- What happens if a user inputs a negative or zero time allocation?
- How does the system handle extremely long task names or descriptions?
- What happens to tasks from previous days? Are they archived, deleted, or carried over?

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST allow users to create a new task with a title.
- **FR-002**: System MUST allow users to allocate time for a task as an estimated duration (e.g., 2 hours or 30 minutes).
- **FR-003**: System MUST track the status of each task with three distinct states: "Not Started", "In Progress", and "Done".
- **FR-004**: System MUST allow users to transition a task between the three states.
- **FR-005**: System MUST allow users to view all tasks grouped or filtered by their current status.
- **FR-006**: System MUST handle daily rollover for tasks. Uncompleted tasks must automatically carry over to the next day.

### Key Entities *(include if feature involves data)*

- **Task**: Represents a single unit of work.
  - Attributes: Title, Status (Not Started, In Progress, Done), Allocated Time, Creation Date.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users can create a new task with a time allocation in under 15 seconds.
- **SC-002**: Users can update the status of a task in a single interaction (e.g., one click or command).
- **SC-003**: System displays the task list and their statuses instantly.

## Assumptions

- The application is for a single user (no multi-user authentication required for MVP).
- Time allocation is initially assumed to be a simple duration (e.g., minutes/hours) rather than complex calendar scheduling.
- The application will be a CLI or simple Web UI (depending on user preference later), but the spec focuses on the abstract functionality.
