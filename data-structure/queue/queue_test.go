package queue

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	tests := []struct {
		name string
		want *Queue
	}{
		// TODO: Add test cases.
		{
			name: "Make a empty queue.",
			want: &Queue{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_EnQueue(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Enqueue 1 to the queue.",
			args: args{
				data: 1,
			},
		},
		{
			name: "Enqueue 2 to the queue.",
			args: args{
				data: 2,
			},
		},
		{
			name: "Enqueue 3 to the queue.",
			args: args{
				data: 3,
			},
		},
		{
			name: "Enqueue 4 to the queue.",
			args: args{
				data: 4,
			},
		},
		{
			name: "Enqueue 5 to the queue.",
			args: args{
				data: 5,
			},
		},
	}
	q := NewQueue()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := q.EnQueue(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Queue.EnQueue() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("After enqueue data %d, Current queue is: %s", tt.args.data, q)
		})
	}
}

func TestQueue_DeQueue(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "dequeue 1.",
			want:    1,
			wantErr: false,
		},
		{
			name:    "dequeue 2.",
			want:    2,
			wantErr: false,
		},
		{
			name:    "dequeue 3.",
			want:    3,
			wantErr: false,
		},
		{
			name:    "dequeue 4.",
			want:    4,
			wantErr: false,
		},
		{
			name:    "dequeue 5.",
			want:    5,
			wantErr: false,
		},
		{
			name:    "dequeue error.",
			want:    0,
			wantErr: true,
		},
	}
	q := NewQueue()
	for i := 1; i < 6; i++ {
		q.EnQueue(i)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := q.DeQueue()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.DeQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queue.DeQueue() = %v, want %v", got, tt.want)
			}
			t.Logf("Got data %d, Curent queue is %s", got, q)
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	q1 := NewQueue()
	q1.EnQueue(1)
	q0 := NewQueue()

	tests := []struct {
		name    string
		queue   *Queue
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Get a head data.",
			queue:   q1,
			want:    1,
			wantErr: false,
		},
		{
			name:    "Get error",
			queue:   q0,
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.queue
			got, err := q.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queue.Peek() = %v, want %v", got, tt.want)
			}
			t.Logf("Got data %d, Current queue is %s", got, q)
		})
	}
}
