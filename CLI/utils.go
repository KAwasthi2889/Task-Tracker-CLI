package main

import (
	"log"
	"strconv"
	"time"
)

func argParser(args []string) (string, []string, bool) {
	if len(args) < 2 {
		log.Println("Usage: task <command> <task>")
		return "", nil, false
	}
	if len(args) < 3 {
		return args[1], []string{"all"}, true
	}
	return args[1], args[2:], true
}

func OutOfBounds(id int, length int) bool {
	if id < 1 || id > length {
		log.Println("Invalid Task ID")
		return true
	}
	return false
}

func is_id(s string) (int, bool) {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Println("InvalId Task number!!")
		return -1, false
	}
	return num, true
}

func add(Description string, tasks_ptr *[]Task) {
	if Description == "" {
		log.Println("Invalid Task!!")
		return
	}
	new_Id := 0
	for _, task := range *tasks_ptr {
		if task.Description == Description {
			log.Println("Following Task already exists:\n", task)
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
	log.Println("Following task added:\n", task)
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
	log.Println("Following task deleted:\n", task)
}

func update(id int, description string, tasks_ptr *[]Task) {
	if OutOfBounds(id, len(*tasks_ptr)) {
		return
	}

	for _, task := range *tasks_ptr {
		if task.Description == description {
			log.Println("Following Task already exists:\n", task)
			return
		}
	}

	task := &(*tasks_ptr)[id-1]
	task.Description = description
	task.UpdateAt = time.Now()
	log.Println("Following task updated:\n", *task)
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
	log.Println("Following task marked as", Status, ":\n", *task)
}

func list(task_type string, tasks_ptr *[]Task) {
	if task_type == "all" {
		for _, task := range *tasks_ptr {
			log.Println(task)
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
		log.Println("Invalid Task Status")
		return
	}
	for _, task := range *tasks_ptr {
		if task.Status == status {
			log.Println(task)
		}
	}
}
