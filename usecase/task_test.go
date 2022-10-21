package usecase

import (
	"reflect"
	"testing"
	"time"

	"github.com/y-mabuchi/go-ddd-todo/domain"

	"github.com/golang/mock/gomock"
	mock_infra "github.com/y-mabuchi/go-ddd-todo/mock/infra"

	"github.com/y-mabuchi/go-ddd-todo/infra"
)

func TestNewTaskUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_infra.NewMockTaskRepositoryInterface(ctrl)

	type args struct {
		taskRepo infra.TaskRepositoryInterface
	}

	tests := []struct {
		name string
		args args
		want *TaskUseCase
	}{
		{
			name: "success",
			args: args{
				taskRepo: mock,
			},
			want: &TaskUseCase{
				repo: mock,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskUseCase(tt.args.taskRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskUseCase_CreateTask(t1 *testing.T) {
	dueDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	success := &domain.Task{
		ID:            0,
		TaskStatus:    domain.UnDone,
		Name:          "task test",
		DueDate:       dueDate,
		PostponeCount: 0,
	}

	type fields struct {
		repo mock_infra.MockTaskRepositoryInterface
	}

	type args struct {
		name    string
		dueDate time.Time
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(m *mock_infra.MockTaskRepositoryInterface)
		want    *domain.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				name:    "task test",
				dueDate: dueDate,
			},
			prepare: func(m *mock_infra.MockTaskRepositoryInterface) {
				m.EXPECT().Save(success).Return(success, nil)
			},
			want:    success,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			ctrl := gomock.NewController(t1)
			defer ctrl.Finish()

			mock := mock_infra.NewMockTaskRepositoryInterface(ctrl)

			tt.prepare(mock)

			t := &TaskUseCase{
				repo: mock,
			}

			got, err := t.CreateTask(tt.args.name, tt.args.dueDate)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("CreateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}
