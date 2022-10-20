package domain

import "time"

type Task struct {
	ID            int
	TaskStatus    TaskStatus
	Name          string
	DueDate       time.Time
	PostponeCount int
}

type TaskStatus string

const (
	UnDone TaskStatus = "undone"
)
