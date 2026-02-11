# 📝 Task Tracker

A simple and efficient command-line task tracker built with Go. This project is a solution to the [task-tracker challenge](https://roadmap.sh/projects/task-tracker) from roadmap.sh.

## ✨ Features

- ➕ Add, update, and delete tasks
- 🏷️ Mark tasks with different statuses (todo, in-progress, done)
- 📋 List tasks by status or view all tasks
- 💾 Persistent storage using JSON
- 🚀 Fast and lightweight CLI

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

Clone the repository:

```bash
git clone git@github.com:maxkaiser11/task-tracker.git
cd backend-projects/task-tracker
```

Build the project:

```bash
go build -o task-tracker
```

## 📖 Usage

### View all available commands

```bash
./task-tracker --help
```

### Managing Tasks

**Add a new task:**
```bash
./task-tracker add "Buy groceries"
```

**Update an existing task:**
```bash
./task-tracker update 1 "Buy groceries and cook dinner"
```

**Delete a task:**
```bash
./task-tracker delete 1
```

### Task Status Management

**Mark a task as in progress:**
```bash
./task-tracker mark-in-progress 1
```

**Mark a task as done:**
```bash
./task-tracker mark-done 1
```

**Mark a task as todo:**
```bash
./task-tracker mark-todo 1
```

### Listing Tasks

**List all tasks:**
```bash
./task-tracker list
```

**List tasks by status:**
```bash
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
```

## 🛠️ Built With

- [Go](https://golang.org/) - The programming language used

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🔗 Links

- [Challenge Source](https://roadmap.sh/projects/task-tracker)
- [Repository](https://github.com/maxkaiser11/task-tracker)

---

Made with ❤️ by [Max](https://github.com/maxkaiser11)