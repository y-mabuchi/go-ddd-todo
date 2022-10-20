package domain_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/y-mabuchi/go-ddd-todo/domain"
)

func TestNewTask(t *testing.T) {
	type args struct {
		name    string
		dueDate time.Time
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				name:    "test",
				dueDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: &domain.Task{
				Name:          "test",
				DueDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				TaskStatus:    domain.UnDone,
				PostponeCount: 0,
			},
			wantErr: false,
		},
		{
			name: "name is empty",
			args: args{
				name:    "",
				dueDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "dueDate is empty",
			args: args{
				name:    "test",
				dueDate: time.Time{},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewTask(tt.args.name, tt.args.dueDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_Postpone(t1 *testing.T) {
	type fields struct {
		ID            int
		TaskStatus    domain.TaskStatus
		Name          string
		DueDate       time.Time
		PostponeCount int
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				ID:            1,
				TaskStatus:    domain.UnDone,
				Name:          "test",
				DueDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				PostponeCount: 0,
			},
			wantErr: false,
		},
		{
			name: "postpone count is over",
			fields: fields{
				ID:            1,
				TaskStatus:    domain.UnDone,
				Name:          "test",
				DueDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				PostponeCount: 3,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &domain.Task{
				ID:            tt.fields.ID,
				TaskStatus:    tt.fields.TaskStatus,
				Name:          tt.fields.Name,
				DueDate:       tt.fields.DueDate,
				PostponeCount: tt.fields.PostponeCount,
			}
			if err := t.Postpone(); (err != nil) != tt.wantErr {
				t1.Errorf("Postpone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
