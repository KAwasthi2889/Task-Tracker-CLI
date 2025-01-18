package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const filename = "tasks.json"

type Task struct {
	Id          int    `json:"ID"`
	Description string `json:"Task"`
	Status      int    `json:"Status"`
	// 0: pending, 1: done, 2: in progress 3: skipped
	CreatedAt time.Time `json:"Creation Date"`
	UpdateAt  time.Time `json:"Update Date"`
}

func (t Task) String() string {
	var Status string
	var symbol rune
	switch t.Status {
	case 0:
		Status = "Pending"
		symbol = ' '
	case 1:
		Status = "Done"
		symbol = '\u2713' //unicode for a tick
	case 2:
		Status = "In Progress"
		symbol = '\u279C' // unicode for an arrow
	case 3:
		Status = "Skipped"
		symbol = '\u2717' // unicode for a cross
	}
	return fmt.Sprintf("[%c] ID: %d | %s | %s | Creation Date: %s | Updated At: %s",
		symbol, t.Id, t.Description, Status,
		t.CreatedAt.Format("15:04:05 02/01/2006"),
		t.UpdateAt.Format("15:04:05 02/01/2006"))
}

func OutOfBound(id int, length int) bool {
	if id < 1 || id > length {
		fmt.Println("Invalid Task ID")
		return true
	}
	return false
}

func add(Description string, tasks_ptr *[]Task) {
	new_Id := 0
	for _, task := range *tasks_ptr {
		if task.Description == Description {
			fmt.Println("Following Task already exists:\n", task)
			return
		}
		new_Id = task.Id
	}
	task := Task{
		Id:          new_Id + 1,
		Description: Description,
		Status:      0,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
	*tasks_ptr = append(*tasks_ptr, task)
	fmt.Println("Following task added:\n", task)
}

func delete(id int, tasks_ptr *[]Task) {
	if OutOfBound(id, len(*tasks_ptr)) {
		return
	}

	task := (*tasks_ptr)[id-1]
	for i := id - 1; i < len(*tasks_ptr)-1; i++ {
		(*tasks_ptr)[i] = (*tasks_ptr)[i+1]
		(*tasks_ptr)[i].Id--
	}
	*tasks_ptr = (*tasks_ptr)[:len(*tasks_ptr)-1]
	fmt.Println("Following task deleted:\n", task)
}

func update(id int, tasks_ptr *[]Task) {
	if OutOfBound(id, len(*tasks_ptr)) {
		return
	}

	task := &(*tasks_ptr)[id-1]
	fmt.Print("Update: ")
	var Description string
	fmt.Scanln(&Description)
	task.Description = Description
	task.UpdateAt = time.Now()
	fmt.Println("Following task updated:\n", *task)
}

func mark(id int, tasks_ptr *[]Task, Status string) {
	if OutOfBound(id, len(*tasks_ptr)) {
		return
	}

	task := &(*tasks_ptr)[id-1]
	switch Status {
	case "pending":
		task.Status = 0
	case "done":
		task.Status = 1
	case "skip":
		task.Status = 3
	case "inprogress":
		task.Status = 2
	}
	task.UpdateAt = time.Now()
	fmt.Println("Following task marked as", Status, ":\n", *task)
}

func list(task_type string, tasks_ptr *[]Task) {
	if task_type == "all" {
		for _, task := range *tasks_ptr {
			fmt.Println(task)
		}
		return
	}
	var status int
	switch task_type {
	case "done":
		status = 1
	case "pending":
		status = 0
	case "inprogress":
		status = 2
	case "skipped":
		status = 3
	default:
		fmt.Println("InvalId Task Status")
		return
	}
	for _, task := range *tasks_ptr {
		if task.Status == status {
			fmt.Println(task)
		}
	}
}

func main() {
	args := len(os.Args)
	if args < 2 {
		fmt.Println("Usage: task-cli <command> <Your Task>")
		return
	}
	command := strings.ToLower(os.Args[1])
	task := ""

	// only list is allowded to have 2 arguments
	if args == 2 {
		if command == "list" {
			task = "all"
		} else {
			return
		}
	}

	if args == 3 {
		task = os.Args[2]
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

	is_num := true
	num, err := strconv.Atoi(task)
	if err != nil {
		is_num = false
	}

	switch command {
	case "add":
		add(task, &tasks)

	case "delete":
		if !is_num {
			fmt.Println("InvalId Task number!!")
			return
		}
		delete(num, &tasks)

	case "update":
		if !is_num {
			fmt.Println("InvalId Task number!!")
			return
		}
		update(num, &tasks)
	case "done", "skip", "inprogress", "pending":
		if !is_num {
			fmt.Println("InvalId Task number!!")
			return
		}
		mark(num, &tasks, command)

	case "list":
		list(strings.ToLower(task), &tasks) // here task is Done, Pending, InProgress, Skipped or All
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
