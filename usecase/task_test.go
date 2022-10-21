package usecase

import (
	"errors"
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
		{
			name: "name is empty",
			args: args{
				name:    "",
				dueDate: dueDate,
			},
			prepare: func(m *mock_infra.MockTaskRepositoryInterface) {
				m.EXPECT().Save(gomock.Any()).AnyTimes()
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
			prepare: func(m *mock_infra.MockTaskRepositoryInterface) {
				m.EXPECT().Save(gomock.Any()).AnyTimes()
			},
			want:    nil,
			wantErr: true,
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

func TestTaskUseCase_PostponeTask(t1 *testing.T) {
	dueDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	success := &domain.Task{
		ID:            1,
		TaskStatus:    domain.UnDone,
		Name:          "task test",
		DueDate:       dueDate,
		PostponeCount: 0,
	}
	postponed := &domain.Task{
		ID:            1,
		TaskStatus:    domain.UnDone,
		Name:          "task test",
		DueDate:       dueDate.AddDate(0, 0, 1),
		PostponeCount: 1,
	}
	maxOver := &domain.Task{
		ID:            1,
		TaskStatus:    domain.UnDone,
		Name:          "task test",
		DueDate:       dueDate,
		PostponeCount: domain.PostponeMaxCount,
	}

	type fields struct {
		repo mock_infra.MockTaskRepositoryInterface
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(m *mock_infra.MockTaskRepositoryInterface)
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			prepare: func(m *mock_infra.MockTaskRepositoryInterface) {
				m.EXPECT().FindById(1).Return(success, nil)
				m.EXPECT().Save(success).Return(postponed, nil)
			},
			wantErr: false,
		},
		{
			name: "postpone count is max",
			args: args{
				id: 1,
			},
			prepare: func(m *mock_infra.MockTaskRepositoryInterface) {
				m.EXPECT().FindById(1).Return(maxOver, nil)
				m.EXPECT().Save(gomock.Any()).AnyTimes()
			},
			wantErr: true,
		},
		{
			name: "not found",
			args: args{
				id: 1,
			},
			prepare: func(m *mock_infra.MockTaskRepositoryInterface) {
				m.EXPECT().FindById(1).Return(nil, errors.New("not found"))
				m.EXPECT().Save(gomock.Any()).AnyTimes()
			},
			wantErr: true,
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

			if err := t.PostponeTask(tt.args.id); (err != nil) != tt.wantErr {
				t1.Errorf("PostponeTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
