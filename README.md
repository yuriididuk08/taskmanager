# CLI Task Manager

A simple command-line interface (CLI) tool to manage daily tasks in Go.  
You can add, delete, view, and mark tasks as completed.

---

## Features

- Add a task (`add`)
- View all tasks (`list`)
- Mark a task as completed (`done`)
- Delete a task (`delete`)
- Tasks are stored in a `tasks.json` file

---

## Project Structure

```

taskmanager/
├── main.go         # CLI entry point
├── tasks.go        # Task management logic
├── tasks.json      # Task storage file
├── go.mod          # Go module file

````

---

## ⚡ Usage

### Add a task

```bash
go run main.go add "Buy groceries"
````

### View all tasks

```bash
go run main.go list
```

Example output:

```text
1. [ ] Buy groceries
2. [ ] Write Go project
```

### Mark a task as completed

```bash
go run main.go done 1
```

### Delete a task

```bash
go run main.go delete 2
```

---

## 💻 Installation

1. Clone the repository:

```bash
git clone https://github.com/yuriididuk08/taskmanager.git
cd taskmanager
```

2. Initialize Go modules (if needed):

```bash
go mod tidy
```

3. Run CLI commands:

```bash
go run main.go [command] [args]
```

Or build a standalone executable:

```bash
go build -o taskmanager main.go tasks.go
./taskmanager [command] [args]
```

---

## 🌟 Examples

```bash
# Add tasks
./taskmanager add "Buy groceries"
./taskmanager add "Write Go project"

# List tasks
./taskmanager list

# Mark a task as done
./taskmanager done 1

# Delete a task
./taskmanager delete 2
```

---

## 📂 Storage

All tasks are stored in the `tasks.json` file in the project root.

---
* Go 1.21+
* JSON for task storage
* CLI implemented with standard `os` and `fmt` packages
