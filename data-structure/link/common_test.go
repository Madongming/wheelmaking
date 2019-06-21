package link

import (
	"reflect"
	"testing"
)

func Test_makeSingleLinkWithArray(t *testing.T) {
	type args struct {
		data []int
	}

	tailNode := &SingleLinkNode{
		Val: 5,
	}
	tests := []struct {
		name string
		args args
		want *SingleLink
	}{
		// TODO: Add test cases.
		{
			name: "Make a Single link.",
			args: args{
				data: []int{1, 2, 3, 4, 5},
			},
			want: &SingleLink{
				size: 5,
				head: &SingleLinkNode{
					Val: 1,
					Next: &SingleLinkNode{
						Val: 2,
						Next: &SingleLinkNode{
							Val: 3,
							Next: &SingleLinkNode{
								Val:  4,
								Next: tailNode,
							},
						},
					},
				},
				tail: tailNode,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeSingleLinkWithArray(tt.args.data)
			t.Logf("Use an array: %v,  make a Link: %s", tt.args.data, got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeSingleLinkWithArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeDoubleLinkWithArray(t *testing.T) {
	type args struct {
		data []int
	}

	node1 := &DoubleLinkNode{
		Val: 1,
	}
	node2 := &DoubleLinkNode{
		Val:  2,
		Prev: node1,
	}
	node1.Next = node2
	node3 := &DoubleLinkNode{
		Val:  3,
		Prev: node2,
	}
	node2.Next = node3
	node4 := &DoubleLinkNode{
		Val:  4,
		Prev: node3,
	}
	node3.Next = node4
	node5 := &DoubleLinkNode{
		Val:  5,
		Prev: node4,
	}
	node4.Next = node5

	tests := []struct {
		name string
		args args
		want *DoubleLink
	}{
		// TODO: Add test cases.
		{
			name: "Make a Single link.",
			args: args{
				data: []int{1, 2, 3, 4, 5},
			},
			want: &DoubleLink{
				size: 5,
				head: node1,
				tail: node5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeDoubleLinkWithArray(tt.args.data)
			t.Logf("Use an array: %v,  make a Link: %s", tt.args.data, got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeDoubleLinkWithArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeCircularSingleLinkWithArray(t *testing.T) {
	type args struct {
		data []int
	}

	node1 := &CircularSingleLinkNode{
		Val: 1,
	}
	node2 := &CircularSingleLinkNode{
		Val: 2,
	}
	node1.Next = node2
	node3 := &CircularSingleLinkNode{
		Val: 3,
	}
	node2.Next = node3
	node4 := &CircularSingleLinkNode{
		Val: 4,
	}
	node3.Next = node4
	node5 := &CircularSingleLinkNode{
		Val: 5,
	}
	node4.Next = node5
	node5.Next = node1

	tests := []struct {
		name string
		args args
		want *CircularSingleLink
	}{
		// TODO: Add test cases.
		{
			name: "Make a Single link.",
			args: args{
				data: []int{1, 2, 3, 4, 5},
			},
			want: &CircularSingleLink{
				size: 5,
				head: node1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeCircularSingleLinkWithArray(tt.args.data)
			t.Logf("Use an array: %v,  make a Link: %s", tt.args.data, got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeSingleLinkWithArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
