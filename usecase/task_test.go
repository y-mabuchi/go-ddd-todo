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

	mockTaskRepo := mock_infra.NewMockTaskRepositoryInterface(ctrl)

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
				taskRepo: mockTaskRepo,
			},
			want: &TaskUseCase{
				repo: mockTaskRepo,
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

func TestTaskUseCase_CreateTask(t *testing.T) {
	task := &domain.Task{
		ID:            0,
		TaskStatus:    domain.UnDone,
		Name:          "task test",
		DueDate:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		PostponeCount: 0,
	}

	type fields struct {
		repo *mock_infra.MockTaskRepositoryInterface
	}

	type args struct {
		name    string
		dueDate time.Time
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(f *fields)
		want    *domain.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				name:    "task test",
				dueDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			prepare: func(f *fields) {
				f.repo.EXPECT().Save(task).Return(task, nil)
			},
			want:    task,
			wantErr: false,
		},
		{
			name: "name is empty",
			args: args{
				name:    "",
				dueDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			prepare: func(f *fields) {
				f.repo.EXPECT().Save(task).Return(task, nil)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "due date is empty",
			args: args{
				name:    "task test",
				dueDate: time.Time{},
			},
			prepare: func(f *fields) {
				f.repo.EXPECT().Save(task).Return(task, nil)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo: mock_infra.NewMockTaskRepositoryInterface(ctrl),
			}

			tt.prepare(&f)

			t := &TaskUseCase{
				repo: tt.fields.repo,
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

func TestTaskUseCase_PostponeTask(t *testing.T) {
	t.Skip("skip")
	task := &domain.Task{
		ID:            0,
		TaskStatus:    domain.UnDone,
		Name:          "task test",
		DueDate:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		PostponeCount: 0,
	}

	// errTask := &domain.Task{
	// 	ID:            1,
	// 	TaskStatus:    domain.UnDone,
	// 	Name:          "task test",
	// 	DueDate:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	PostponeCount: 3,
	// }

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockTaskRepo := mock_infra.NewMockTaskRepositoryInterface(mockCtrl)

	type fields struct {
		repo infra.TaskRepositoryInterface
	}

	type args struct {
		id int
	}

	tests := []struct {
		name          string
		fields        fields
		args          args
		prepareMockFn func(m *mock_infra.MockTaskRepositoryInterface)
		wantErr       bool
	}{
		{
			name: "success",
			fields: fields{
				repo: mockTaskRepo,
			},
			args: args{
				id: 0,
			},
			prepareMockFn: func(m *mock_infra.MockTaskRepositoryInterface) {
				m.EXPECT().FindById(0).Return(task, nil)
				m.EXPECT().Save(task).Return(task, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			ctrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mock := mock_infra.NewMockTaskRepositoryInterface(ctrl)

			tt.prepareMockFn(mock)

			t := &TaskUseCase{
				repo: tt.fields.repo,
			}

			if err := t.PostponeTask(tt.args.id); (err != nil) != tt.wantErr {
				t1.Errorf("PostponeTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
