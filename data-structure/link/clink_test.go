package link

import (
	"reflect"
	"testing"
)

func TestCircularSingleLink_InsertNext(t *testing.T) {
	type args struct {
		node *CircularSingleLinkNode
		data int
	}
	csl := makeCircularSingleLinkWithArray([]int{1, 3, 4, 5})
	var addNode *CircularSingleLinkNode
	dummy := csl.head
	for dummy != nil {
		if dummy.Val == 1 {
			addNode = dummy
			break
		}
		dummy = dummy.Next
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Circular link insert.",
			args: args{
				node: addNode,
				data: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := csl.InsertNext(tt.args.node, tt.args.data)
			if err != nil {
				t.Errorf("CircularSingleLink.InsertNext() error = %v", err)
			}
			want := makeCircularSingleLinkWithArray([]int{1, 2, 3, 4, 5})
			if !reflect.DeepEqual(csl, want) {
				t.Errorf("Got link %v, Want link %v", csl.head, want.head)
			}

			t.Logf("Got link %v, Want link %v", csl.head, want.head)
		})
	}
}

func TestCircularSingleLink_RemoveNext(t *testing.T) {
	type args struct {
		node *CircularSingleLinkNode
	}
	csl := makeCircularSingleLinkWithArray([]int{1, 2, 3, 4, 5})
	var delNode *CircularSingleLinkNode
	dummy := csl.head
	for dummy != nil {
		if dummy.Val == 3 {
			delNode = dummy
			break
		}
		dummy = dummy.Next
	}

	tests := []struct {
		name   string
		fields *CircularSingleLink
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name:   "Remove a node.",
			fields: csl,
			args: args{
				node: delNode,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csl := tt.fields
			t.Logf("Old Link: %s\n", csl.head)
			got, _ := csl.RemoveNext(tt.args.node)
			if got != tt.want {
				t.Errorf("CircularSingleLink.RemoveNext() = %v, want %v", got, tt.want)
			}
			t.Logf("Old Link: %s\n", csl.head)
		})
	}
}
