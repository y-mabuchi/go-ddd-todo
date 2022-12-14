//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package infra

import (
	"log"
	"time"

	"github.com/y-mabuchi/go-ddd-todo/domain"
)

type TaskRepositoryInterface interface {
	Save(task *domain.Task) (*domain.Task, error)
	FindById(id int) (*domain.Task, error)
}

type TaskRepository struct {
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (t *TaskRepository) Save(task *domain.Task) (*domain.Task, error) {
	log.Println("task saved.")

	return task, nil
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
