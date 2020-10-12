package queue

import (
	"reflect"
	"testing"
)

func Test_queue_Empty(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			if got := q.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_PopFront(t *testing.T) {
	tests := []struct {
		name    string
		want    interface{}
		wantErr bool
	}{
		{
			name:    "",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			got, err := q.PopFront()
			if (err != nil) != tt.wantErr {
				t.Errorf("PopFront() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopFront() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_PushBack(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{v: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			q.PushBack(tt.args.v)
			if q.Size() != 1 {
				t.Errorf("Size() got = %v, want = %v", q.Size(), 1)
			}
		})
	}
}

func Test_queue_Size(t *testing.T) {
	tests := []struct {
		name   string
		want   int
	}{
		{
			name: "",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			if got := q.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContinuousPushAndPop(t *testing.T) {
	n := 10000000
	q := NewQueue()
	for i := 0; i < n; i++ {
		q.PushBack(i)
	}
	for i := 0; i < n; i++ {
		if front, err := q.PopFront(); err == nil {
			if i != front {
				t.Errorf("PopFront() got = %v, want = %v", front, i)
				return
			}
		}
	}
}
