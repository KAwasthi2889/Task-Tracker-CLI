package main

import (
	"fmt"
	"time"
)

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
