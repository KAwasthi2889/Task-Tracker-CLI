# CLI Task Tracker

A lightweight command-line interface (CLI) application to manage your tasks efficiently. This app allows you to add, delete, update, and manage the status of your tasks, all while storing them in a JSON file in the current directory.

---

## Features

- **Add Task**: Create new tasks with ease.
- **Delete Task**: Remove tasks you no longer need.
- **Update Task**: Modify task details as needed.
- **Manage Task Status**: Mark tasks as:
  - Done
  - Pending
  - Skipped
  - In-Progress
- **List Tasks**: List tasks as:
  - All
  - Done
  - Pending
  - Skipped
  - In-Progress
- **Data Persistence**: All tasks are stored in a `tasks.json` file in the current directory for easy access and management.

---

## Installation

1. Clone this repository:
   ```bash
   git clone <repository_url>
   ```

2. Navigate to the project directory:
   ```bash
   cd cli-task-tracker
   ```

3. Build the CLI application:
   ```bash
   go build -o task-cli
   ```

---

## Usage

### Adding a Task
```bash
./task-cli add "Task Description"
```

### Deleting a Task
```bash
./task-cli delete <task_id>
```

### Updating a Task
```bash
./task-cli update <task_id> "New Task Description"
```

### Changing Task Status
- **Mark as Done:**
  ```bash
  ./task-cli done <task_id>
  ```
- **Mark as Pending:**
  ```bash
  ./task-cli pending <task_id>
  ```
- **Mark as Skipped:**
  ```bash
  ./task-cli skipped <task_id>
  ```
- **Mark as In-Progress:**
  ```bash
  ./task-cli in-progress <task_id>
  ```

---

### Listing Tasks
- **Listing all Tasks:**
  ```bash
  ./task-cli list
  ```
- **Listing done Tasks:**
  ```bash
  ./task-cli list done
  ```
- **Listing pending Tasks:**
  ```bash
  ./task-cli list pending
  ```
- **Listing in-progress Tasks:**
  ```bash
  ./task-cli list in-progress
  ```
- **Listing skipped Tasks:**
  ```bash
  ./task-cli list skipped
  ```
 
---

## JSON File Structure

Tasks are stored in a `tasks.json` file in the current directory. Example structure:

```json
[
  {
    "id": 1,
    "description": "Example Task",
    "status": 0,
    "Creation Date": "2025-01-19T15:14:25.958703952+05:30",
    "Update Date": "2025-01-19T15:14:25.958704095+05:30"
  },
  {
    "id": 2,
    "description": "Another Task",
    "status": 1,
    "Creation Date": "2025-01-19T15:14:25.958703952+05:30",
    "Update Date": "2025-01-19T15:14:25.958704095+05:30"
  }
]
```

---

## Contribution

Feel free to fork this repository and submit pull requests.
Any contributions are welcome to improve functionality or add new features.

---

## License

This project is licensed under the [MIT License](LICENSE).