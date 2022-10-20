package infra

import (
	"github.com/y-mabuchi/go-ddd-todo/domain"
	"log"
	"time"
)

type TaskRepositoryInterface interface {
	Save(task *domain.Task) error
	FindById(id int) (*domain.Task, error)
}

type TaskRepository struct {
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (t *TaskRepository) Save(task *domain.Task) error {
	log.Println("task saved.")

	return nil
}

func (t *TaskRepository) FindById(id int) (*domain.Task, error) {
	task := &domain.Task{
		ID:            id,
		TaskStatus:    domain.UnDone,
		Name:          "task test",
		DueDate:       time.Now(),
		PostponeCount: 0,
	}

	return task, nil
}
