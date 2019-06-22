package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want *Stack
	}{
		// TODO: Add test cases.
		{
			name: "Make a empty stack.",
			want: &Stack{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
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
			name: "Push 1 to the stack.",
			args: args{
				data: 1,
			},
		},
		{
			name: "Push 2 to the stack.",
			args: args{
				data: 2,
			},
		},
		{
			name: "Push 3 to the stack.",
			args: args{
				data: 3,
			},
		},
		{
			name: "Push 4 to the stack.",
			args: args{
				data: 4,
			},
		},
		{
			name: "Push 5 to the stack.",
			args: args{
				data: 5,
			},
		},
	}
	s := NewStack()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Push(tt.args.data); err != nil {
				t.Errorf("Stack.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("After push data %d, Current stack is: %s", tt.args.data, s)
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Pop 5.",
			want:    5,
			wantErr: false,
		},
		{
			name:    "Pop 4.",
			want:    4,
			wantErr: false,
		},
		{
			name:    "Pop 3.",
			want:    3,
			wantErr: false,
		},
		{
			name:    "Pop 2.",
			want:    2,
			wantErr: false,
		},
		{
			name:    "Pop 1.",
			want:    1,
			wantErr: false,
		},
		{
			name:    "Pop error.",
			want:    0,
			wantErr: true,
		},
	}

	s := NewStack()
	for i := 1; i < 6; i++ {
		s.Push(i)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
			t.Logf("Got data %d, Curent stack is %s", got, s)
		})
	}
}

func TestStack_Peek(t *testing.T) {
	s1 := NewStack()
	s1.Push(1)
	s0 := NewStack()

	tests := []struct {
		name    string
		stack   *Stack
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Get a top data.",
			stack:   s1,
			want:    1,
			wantErr: false,
		},
		{
			name:    "Get error",
			stack:   s0,
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.stack
			got, err := s.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Stack.Peek() = %v, want %v", got, tt.want)
			}
			t.Logf("Got data %d, Current stack is %s", got, s)
		})
	}
}
