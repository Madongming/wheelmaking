package link

import (
	"reflect"
	"testing"
)

func TestDoubleLink_InsertNext(t *testing.T) {
	type args struct {
		node *DoubleLinkNode
		data int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Dlink insert next.",
			args: args{
				node: &DoubleLinkNode{
					Val: 2,
				},
				data: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := makeDoubleLinkWithArray([]int{1, 2, 4, 5})
			err := dl.InsertNext(tt.args.node, tt.args.data)
			if err != nil {
				t.Errorf("DoubleLink.InsertNext() error = %v.", err)
			}
			want := makeDoubleLinkWithArray([]int{1, 2, 3, 4, 5})
			if !reflect.DeepEqual(dl, want) {
				t.Errorf("Got double link %v, want double link: %s", dl, want)
			}
			t.Logf("Got double link %v, want double link: %s", dl, want)
		})
	}
}

func TestDoubleLink_InsertPrev(t *testing.T) {
	type args struct {
		node *DoubleLinkNode
		data int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Dlink insert next.",
			args: args{
				node: &DoubleLinkNode{
					Val: 3,
				},
				data: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := makeDoubleLinkWithArray([]int{1, 3, 4, 5})
			err := dl.InsertPrev(tt.args.node, tt.args.data)
			if err != nil {
				t.Errorf("DoubleLink.InsertNext() error = %v", err)
			}
			want := makeDoubleLinkWithArray([]int{1, 2, 3, 4, 5})
			if !reflect.DeepEqual(dl, want) {
				t.Errorf("Got double link %v, want double link: %s", dl, want)
			}
			t.Logf("Got double link %v, want double link: %s", dl, want)
		})
	}
}

func TestDoubleLink_Remove(t *testing.T) {
	type args struct {
		node *DoubleLinkNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Delete a node.",
			args: args{
				node: &DoubleLinkNode{
					Val: 3,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := makeDoubleLinkWithArray([]int{1, 2, 3, 4, 5})
			t.Logf("Old double link: %s", dl.head)
			got, err := dl.Remove(tt.args.node)
			if err != nil {
				t.Errorf("DoubleLink.Remove() error = %v.", err)
				return
			}
			if got != tt.want {
				t.Errorf("DoubleLink.Remove() = %v, want %v", got, tt.want)
			}
			t.Logf("New double link: %s", dl.head)
		})
	}
}
