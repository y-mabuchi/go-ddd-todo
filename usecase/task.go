package usecase

import (
	"time"

	"github.com/y-mabuchi/go-ddd-todo/domain"
	"github.com/y-mabuchi/go-ddd-todo/infra"
)

type TaskUseCaseInterface interface {
	CreateTask(name string, dueDate time.Time) error
	PostponeTask(id int) error
}

type TaskUseCase struct {
	repo infra.TaskRepositoryInterface
}

func NewTaskUseCase(taskRepo infra.TaskRepositoryInterface) *TaskUseCase {
	return &TaskUseCase{
		repo: taskRepo,
	}
}

func (t *TaskUseCase) CreateTask(name string, dueDate time.Time) error {
	task, err := domain.NewTask(name, dueDate)
	if err != nil {
		return err
	}

	if err := t.repo.Save(task); err != nil {
		return err
	}

	return nil
}

func (t *TaskUseCase) PostponeTask(id int) error {
	task, err := t.repo.FindById(id)
	if err != nil {
		return err
	}

	if err = task.Postpone(); err != nil {
		return err
	}

	if err = t.repo.Save(task); err != nil {
		return err
	}

	return nil
}
