package main

import (
	"bytes"
	"log"
	"testing"
)

var buf bytes.Buffer

func Test_argParser(t *testing.T) {
	log.SetOutput(&buf)

	arguments := [][]string{
		{"main.go", "add", "task1"},
		{"main.go", "delete", "1"},
		{"main.go", "update", "1", "task2"},
		{"main.go", "mark", "1", "done"},
		{"main.go", "list"},
		{"main.go", "list", "done"},
	}

	expected_commands := []string{
		"add",
		"delete",
		"update",
		"mark",
		"list",
		"list",
	}

	expected_tasks := [][]string{
		{"task1"},
		{"1"},
		{"1", "task2"},
		{"1", "done"},
		{"all"},
		{"done"},
	}
	for i, args := range arguments {
		command, tasks, ok := argParser(args)
		if command != expected_commands[i] {
			t.Errorf("Expected %s, got %s", expected_commands[i], command)
		}
		for j, task := range tasks {
			if task != expected_tasks[i][j] {
				t.Errorf("Expected %s, got %s", expected_tasks[i][j], task)
			}
		}
		if ok != true {
			t.Errorf("Expected true, got %t", ok)
		}
	}
}

func Test_OutOfBounds(t *testing.T) {
	log.SetOutput(&buf)

	arguments := [][]int{
		{5, 4},
		{0, 1},
		{1, 1},
	}

	expected := []bool{true, true, false}

	for i, args := range arguments {
		if OutOfBounds(args[0], args[1]) != expected[i] {
			t.Errorf("Expected %t, got %t", expected[i], !expected[i])
		}
	}
}

func Test_is_id(t *testing.T) {
	log.SetOutput(&buf)

	arguments := []string{"5", "0", "a"}

	expected := []bool{true, true, false}

	for i, arg := range arguments {
		_, ok := is_id(arg)
		if ok != expected[i] {
			t.Errorf("Expected %t, got %t", expected[i], !expected[i])
		}
	}
}

func Test_add(t *testing.T) {
	log.SetOutput(&buf)

	arguments := []string{"task1", "", "task 1 and 2"}

	var tasks []Task

	for _, arg := range arguments {
		add(arg, &tasks)
	}

	expected_tasks := []Task{
		{Id: 1, Description: "task1", Status: 0},
		{Id: 2, Description: "task 1 and 2", Status: 0},
	}

	for i, task := range tasks {
		if task.Id != expected_tasks[i].Id {
			t.Errorf("Expected %d, got %d", expected_tasks[i].Id, task.Id)
		}
		if task.Description != expected_tasks[i].Description {
			t.Errorf("Expected %s, got %s", expected_tasks[i].Description, task.Description)
		}
		if task.Status != expected_tasks[i].Status {
			t.Errorf("Expected %d, got %d", expected_tasks[i].Status, task.Status)
		}
	}
}

func Test_delete(t *testing.T) {
	log.SetOutput(&buf)

	tasks := []Task{
		{Id: 1, Description: "task1"},
		{Id: 2, Description: "task2"},
		{Id: 3, Description: "task3"},
		{Id: 4, Description: "task4"},
	}

	arguments := []int{0, 2, 4}

	for _, arg := range arguments {
		delete(arg, &tasks)
	}

	if task := tasks[1]; task.Id != 2 || task.Description != "task3" {
		t.Errorf("Expected 2, got %d", task.Id)
		t.Errorf("Expected task3, got %s", task.Description)
	}

}

func Test_update(t *testing.T) {
	log.SetOutput(&buf)

	tasks := []Task{
		{Id: 1, Description: "task1"},
		{Id: 2, Description: "task2"},
		{Id: 3, Description: "task3"},
		{Id: 4, Description: "task4"},
	}

	arguments := []string{
		"New task 1",
		"New task 2",
		"New task 3",
		"New task 4",
	}

	for i, arg := range arguments {
		update(i+1, arg, &tasks)
	}

	for i, task := range tasks {
		if task.Description != arguments[i] {
			t.Errorf("Expected %s, got %s", arguments[i], task.Description)
		}
	}
}

func Test_mark(t *testing.T) {
	log.SetOutput(&buf)

	tasks := []Task{
		{Id: 1, Description: "task1", Status: 0},
		{Id: 2, Description: "task2", Status: 0},
		{Id: 3, Description: "task3", Status: 0},
		{Id: 4, Description: "task4", Status: 0},
	}

	arguments := []string{"done", "in-progress", "skipped", "pending"}
	expected := []int{1, 2, 3, 0}

	for i, status := range arguments {
		mark(i+1, &tasks, status)
	}

	for i, task := range tasks {
		if task.Status != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], task.Status)
		}
	}
}
