package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const filename = "tasks.json"

type Task struct {
	id          int    `json:"ID"`
	description string `json:"Task"`
	status      int    `json:"Status"`
	// 0: pending, 1: done, 2: in progress 3: skipped
	createdAt time.Time `json:"Creation Date"`
	updateAt  time.Time `json:"Update Date"`
}

func add(description string, tasks []Task) []Task {
	task := Task{
		id:          tasks[len(tasks)-1].id + 1,
		description: description,
		status:      0,
		createdAt:   time.Now(),
		updateAt:    time.Now(),
	}
	tasks = append(tasks, task)
	fmt.Println("Task added successfully")
	return tasks
}

func delete(index int, tasks []Task) []Task {
	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Println("Task deleted successfully")
	return tasks
}

func main() {
	args := len(os.Args)

	if args != 3 {
		if args == 2 && os.Args[1] == "List" {
			list()
			return
		}
		fmt.Println("Usage: task-cli <command> \"<Your Task>\"")
		return
	}

	command := os.Args[1]
	description := os.Args[2]

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755) // this number is file permission
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var tasks []Task
	json.Unmarshal(data, &tasks)

	present := false
	index := -1
	for i, t := range tasks {
		if t.description == description {
			present = true
			index = i
			break
		}
	}

	switch command {
	case "Add":
		if present {
			fmt.Println("Task already exists")
		} else {
			tasks = add(description, tasks)
		}
	case "Delete":
		if !present {
			fmt.Println("Task not found")
		} else {
			tasks = delete(index, tasks)
		}
	case "Update":
		if !present {
			tasks = add(description, tasks)
		} else {
			tasks = update(index, tasks)
		}
	case "List":
		list(task)
	case "Done":
		done(task)
	case "Skip":
		skip(task)
	case "InProgress":
		inProgress(task)
	}
}
