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

func write_file(args []string) {
	var todos []Task

	// Read whole file
	data, err := os.ReadFile("db.json")
	if err != nil {
		// If file doesn't exist yet, start with empty list
		if os.IsNotExist(err) {
			todos = []Task{}
		} else {
			panic(err)
		}
	} else {
		// If file is empty or whitespace, treat as empty JSON array
		if len(bytes.TrimSpace(data)) == 0 {
			todos = []Task{}
		} else if err := json.Unmarshal(data, &todos); err != nil {
			panic(err)
		}
	}

	if args[1] == "add" {
		fmt.Println("Adding task...")

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
	}

	if args[1] == "update" {
		id, _ := strconv.Atoi(args[2])
		for i := range todos {
			if todos[i].ID == id {
				todos[i].Description = args[3]
			}
		}
	}

	if args[1] == "delete" {
		for i := range todos {
			if strconv.Itoa(todos[i].ID) == args[2] {
				todos = append(todos[:i], todos[i+1:]...)
				fmt.Printf("Task deleted")
				break
			}
		}
	}

	if args[1] == "mark-in-progress" {
		for i := range todos {
			if strconv.Itoa(todos[i].ID) == args[2] {
				todos[i].Status = "In Progress"
				break
			}
		}
		fmt.Printf("Task In Progress")
	}

	if args[1] == "mark-done" {
		for i := range todos {
			if strconv.Itoa(todos[i].ID) == args[2] {
				todos[i].Status = "Done"
				break
			}
		}
		fmt.Printf("Task Completed")
	}

	if len(args) == 3 && args[1] == "list" && args[2] == "done" {
		var doneTasks []Task
		for i := range todos {
			if todos[i].Status == "Done" {
				doneTasks = append(doneTasks, todos[i])
			}
		}
		if len(doneTasks) == 0 {
			println("No Tasks marked as Done")
		}
		for i := range doneTasks {
			fmt.Printf("ID: %d, Description: %s\n", doneTasks[i].ID, doneTasks[i].Description)
		}
	}

	if len(args) == 3 && args[1] == "list" && args[2] == "todo" {
		var doneTasks []Task
		for i := range todos {
			if todos[i].Status == "todo" {
				doneTasks = append(doneTasks, todos[i])
			}
		}
		if len(doneTasks) == 0 {
			println("No Tasks marked as Todo")
		}

		for i := range doneTasks {
			fmt.Printf("ID: %d, Description: %s\n", doneTasks[i].ID, doneTasks[i].Description)
		}
	}

	if len(args) == 3 && args[1] == "list" && args[2] == "in-progress" {
		var doneTasks []Task
		for i := range todos {
			if todos[i].Status == "In Progress" {
				doneTasks = append(doneTasks, todos[i])
			}
		}
		if len(doneTasks) == 0 {
			println("No Tasks marked as In Progress")
		}
		for i := range doneTasks {
			fmt.Printf("ID: %d, Description: %s\n", doneTasks[i].ID, doneTasks[i].Description)
		}
	}

	if len(args) == 2 && args[1] == "list" {
		for _, todo := range todos {
			fmt.Printf("ID: %d, Description: %s, Status: %s\n", todo.ID, todo.Description, todo.Status)
		}
	}

	updated, _ := json.MarshalIndent(todos, "", " ")
	os.WriteFile("db.json", updated, 0644)
}

func main() {
	println("Please choose from the following options: ")
	println("1. List Tasks (list) \n2. Add Task (add) \n3. Update Task (update) \n4. Mark task Complete (complete) \n5. Delete Task (delete)")

	args := os.Args

	write_file(args)

}
