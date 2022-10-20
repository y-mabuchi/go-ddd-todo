package usecase

import (
	"errors"
	"github.com/y-mabuchi/go-ddd-todo/domain"
	"github.com/y-mabuchi/go-ddd-todo/infrastructure"
	"time"
)

type TaskUseCaseInterface interface {
	CreateTask(name string, dueDate time.Time) error
	PostponeTask(id int) error
}

type TaskUseCase struct {
	taskRepo         infrastructure.TaskRepositoryInterface
	postponeMaxCount int
}

func NewTaskUseCase(taskRepo infrastructure.TaskRepositoryInterface, postponeMaxCount int) *TaskUseCase {
	return &TaskUseCase{
		taskRepo:         taskRepo,
		postponeMaxCount: postponeMaxCount,
	}
}

func (t *TaskUseCase) CreateTask(name string, dueDate time.Time) error {
	if name == "" || dueDate.IsZero() {
		return errors.New("name or dueDate is empty")
	}

	task := &domain.Task{
		TaskStatus:    domain.UnDone,
		Name:          name,
		DueDate:       dueDate,
		PostponeCount: 0,
	}

	if err := t.taskRepo.Save(task); err != nil {
		return err
	}

	return nil
}

func (t *TaskUseCase) PostponeTask(id int) error {
	task, err := t.taskRepo.FindById(id)
	if err != nil {
		return err
	}

	if task.PostponeCount >= t.postponeMaxCount {
		return errors.New("postpone count is over")
	}

	task.DueDate = task.DueDate.AddDate(0, 0, 1)
	task.PostponeCount++

	if err = t.taskRepo.Save(task); err != nil {
		return err
	}

	return nil
}
