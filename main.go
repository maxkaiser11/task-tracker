package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func showHelp() {
	fmt.Println("Task Tracker - A simple CLI task management tool")
	fmt.Println("\nUsage: task-tracker [command] [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  add [description]           - Add a new task")
	fmt.Println("  update [id] [description]   - Update a task")
	fmt.Println("  delete [id]                 - Delete a task")
	fmt.Println("  mark-in-progress [id]       - Mark task as in progress")
	fmt.Println("  mark-done [id]              - Mark task as done")
	fmt.Println("  list                        - List all tasks")
	fmt.Println("  list done                   - List completed tasks")
	fmt.Println("  list todo                   - List pending tasks")
	fmt.Println("  list in-progress            - List tasks in progress")
	fmt.Println("\nExamples:")
	fmt.Println("  ./task-tracker add \"Buy groceries\"")
	fmt.Println("  ./task-tracker update 1 \"Buy groceries and cook dinner\"")
	fmt.Println("  ./task-tracker mark-done 1")
	fmt.Println("  ./task-tracker list done")
}

func loadTasks() []Task {
	var todos []Task
	data, err := os.ReadFile("db.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}
		}
		panic(err)
	}

	if len(bytes.TrimSpace(data)) == 0 {
		return []Task{}
	}

	if err := json.Unmarshal(data, &todos); err != nil {
		panic(err)
	}

	return todos
}

func saveTasks(todos []Task) {
	updated, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("db.json", updated, 0644)
}

func handleCommand(args []string) {
	if len(args) < 2 {
		showHelp()
		return
	}

	todos := loadTasks()
	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Error: Please provide a task description")
			return
		}

		nextID := 1
		if len(todos) > 0 {
			nextID = todos[len(todos)-1].ID + 1
		}

		task := Task{
			ID:          nextID,
			Description: args[2],
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		todos = append(todos, task)
		fmt.Printf("Task added successfully (ID: %d)\n", task.ID)

	case "update":
		if len(args) < 4 {
			fmt.Println("Error: Please provide task ID and new description")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID")
			return
		}

		found := false
		for i := range todos {
			if todos[i].ID == id {
				todos[i].Description = args[3]
				todos[i].UpdatedAt = time.Now()
				found = true
				fmt.Printf("Task %d updated successfully\n", id)
				break
			}
		}
		if !found {
			fmt.Printf("Error: Task with ID %d not found\n", id)
		}

	case "delete":
		if len(args) < 3 {
			fmt.Println("Error: Please provide task ID")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID")
			return
		}

		found := false
		for i := range todos {
			if todos[i].ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				fmt.Printf("Task %d deleted successfully\n", id)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Error: Task with ID %d not found\n", id)
		}

	case "mark-in-progress":
		if len(args) < 3 {
			fmt.Println("Error: Please provide task ID")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID")
			return
		}

		found := false
		for i := range todos {
			if todos[i].ID == id {
				todos[i].Status = "in-progress"
				todos[i].UpdatedAt = time.Now()
				fmt.Printf("Task %d marked as in progress\n", id)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Error: Task with ID %d not found\n", id)
		}

	case "mark-done":
		if len(args) < 3 {
			fmt.Println("Error: Please provide task ID")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID")
			return
		}

		found := false
		for i := range todos {
			if todos[i].ID == id {
				todos[i].Status = "done"
				todos[i].UpdatedAt = time.Now()
				fmt.Printf("Task %d marked as done\n", id)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Error: Task with ID %d not found\n", id)
		}

	case "list":
		if len(args) == 2 {
			// List all tasks
			if len(todos) == 0 {
				fmt.Println("No tasks found")
				return
			}
			for _, todo := range todos {
				fmt.Printf("ID: %d, Description: %s, Status: %s\n", todo.ID, todo.Description, todo.Status)
			}
		} else if len(args) == 3 {
			// List by status
			status := args[2]
			var filteredTasks []Task

			switch status {
			case "done":
				for _, todo := range todos {
					if todo.Status == "done" {
						filteredTasks = append(filteredTasks, todo)
					}
				}
			case "todo":
				for _, todo := range todos {
					if todo.Status == "todo" {
						filteredTasks = append(filteredTasks, todo)
					}
				}
			case "in-progress":
				for _, todo := range todos {
					if todo.Status == "in-progress" {
						filteredTasks = append(filteredTasks, todo)
					}
				}
			default:
				fmt.Printf("Error: Unknown status '%s'\n", status)
				return
			}

			if len(filteredTasks) == 0 {
				fmt.Printf("No tasks with status '%s'\n", status)
				return
			}

			for _, todo := range filteredTasks {
				fmt.Printf("ID: %d, Description: %s\n", todo.ID, todo.Description)
			}
		}

	default:
		fmt.Printf("Error: Unknown command '%s'\n", command)
		fmt.Println("Run './task-tracker --help' for usage information")
		return
	}

	saveTasks(todos)
}

func main() {
	args := os.Args

	// Check for help flag
	if len(args) < 2 || args[1] == "--help" || args[1] == "-h" {
		showHelp()
		return
	}

	handleCommand(args)
}
