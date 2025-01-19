package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const filename = "tasks.json"

func main() {
	command, task, ok := argParser(os.Args)
	if !ok {
		return
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
		add(strings.Join(task, " "), &tasks)

	case "update":
		if num, ok := is_id(task[0]); !ok {
			return
		} else {
			update(num, strings.Join(task[1:], " "), &tasks)
		}

	case "delete":
		if num, ok := is_id(task[0]); !ok {
			return
		} else {
			delete(num, &tasks)
		}

	case "done", "skipped", "in-progress", "pending":
		if num, ok := is_id(task[0]); !ok {
			return
		} else {
			mark(num, &tasks, command)
		}

	case "list":
		list(strings.ToLower(task[0]), &tasks)
		// here task is Done, Pending, In-Progress, Skipped or All

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
