package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const filename = "tasks.json"

func main() {
	args := len(os.Args)
	if args < 2 {
		fmt.Println("Usage: task-cli <command> <Your Task/ Task id>")
		return
	}
	command := strings.ToLower(os.Args[1])
	task := ""

	if args == 2 {
		if command == "list" {
			task = "all"
		} else {
			fmt.Println("Usage: task-cli <command> <Your Task/ Task id>")
			return
		}
	} else {
		task = strings.Join(os.Args[2:], " ")
	}

	// Open the file in read-write mode and create it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755) // this number is file permission
	if err != nil {
		fmt.Println("Error Opening File:", err)
		return
	}
	defer file.Close()

	// Read the file and turn the data to []byte
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var tasks []Task
	json.Unmarshal(data, &tasks)

	switch command {
	case "add":
		task = strings.TrimSpace(task)
		add(task, &tasks)

	case "update":
		split := strings.Split(task, " ")
		if num, ok := is_id(split[0]); !ok {
			return
		} else {
			task = strings.Join(split[1:], " ")
			update(num, task, &tasks)
		}

	case "delete":
		if num, ok := is_id(task); !ok {
			return
		} else {
			delete(num, &tasks)
		}

	case "done", "skip", "in progress", "pending":
		if num, ok := is_id(task); !ok {
			return
		} else {
			mark(num, &tasks, command)
		}

	case "list":
		list(strings.ToLower(task), &tasks)
		// here task is Done, Pending, In Progress, Skipped or All

	default:
		fmt.Println("Invalid Command!!")
	}

	// Convert the data back to the json
	data, err = json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error converting data to json:", err)
		return
	}

	// Save the data to the file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
