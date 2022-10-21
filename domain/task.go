//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package domain

import (
	"errors"
	"time"
)

const (
	PostponeMaxCount int = 3
)

type TaskInterface interface {
	Postpone() error
}

type Task struct {
	ID            int
	TaskStatus    TaskStatus
	Name          string
	DueDate       time.Time
	PostponeCount int
}

func NewTask(name string, dueDate time.Time) (*Task, error) {
	if name == "" || dueDate.IsZero() {
		return nil, errors.New("name or dueDate is empty")
	}

	return &Task{
		Name:          name,
		DueDate:       dueDate,
		TaskStatus:    UnDone,
		PostponeCount: 0,
	}, nil
}

func (t *Task) Postpone() error {
	if t.PostponeCount >= PostponeMaxCount {
		return errors.New("postpone count is over")
	}

	t.DueDate = t.DueDate.AddDate(0, 0, 1)
	t.PostponeCount++

	return nil
}

type TaskStatus string

const (
	UnDone TaskStatus = "UNDONE"
)
