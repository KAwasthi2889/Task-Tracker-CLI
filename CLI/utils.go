package main

import (
	"fmt"
	"strconv"
	"time"
)

func OutOfBounds(id int, length int) bool {
	if id < 1 || id > length {
		fmt.Println("Invalid Task ID")
		return true
	}
	return false
}

func is_id(s string) (int, bool) {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("InvalId Task number!!")
		return -1, false
	}
	return num, true
}

func add(Description string, tasks_ptr *[]Task) {
	if Description == "" {
		fmt.Println("Invalid Task!!")
		return
	}
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
	if OutOfBounds(id, len(*tasks_ptr)) {
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

func update(id int, description string, tasks_ptr *[]Task) {
	if OutOfBounds(id, len(*tasks_ptr)) {
		return
	}

	for _, task := range *tasks_ptr {
		if task.Description == description {
			fmt.Println("Following Task already exists:\n", task)
			return
		}
	}

	task := &(*tasks_ptr)[id-1]
	task.Description = description
	task.UpdateAt = time.Now()
	fmt.Println("Following task updated:\n", *task)
}

func mark(id int, tasks_ptr *[]Task, Status string) {
	if OutOfBounds(id, len(*tasks_ptr)) {
		return
	}

	task := &(*tasks_ptr)[id-1]
	switch Status {
	case "pending":
		task.Status = 0
	case "done":
		task.Status = 1
	case "skipped":
		task.Status = 3
	case "in-progress":
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
	case "in-progress":
		status = 2
	case "skipped":
		status = 3
	default:
		fmt.Println("Invalid Task Status")
		return
	}
	for _, task := range *tasks_ptr {
		if task.Status == status {
			fmt.Println(task)
		}
	}
}
