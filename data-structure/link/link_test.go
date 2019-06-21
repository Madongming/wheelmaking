package link

import (
	"reflect"
	"testing"
)

func TestSingleLink_InsertNext(t *testing.T) {
	type args struct {
		node *SingleLinkNode
		data int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "link insert.",
			args: args{
				node: &SingleLinkNode{
					Val: 1,
				},
				data: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := makeSingleLinkWithArray([]int{1, 3, 4, 5})
			err := sl.InsertNext(tt.args.node, tt.args.data)
			if err != nil {
				t.Errorf("SingleLink.InsertNext() error = %v", err)
			}
			want := makeSingleLinkWithArray([]int{1, 2, 3, 4, 5})
			if !reflect.DeepEqual(sl, want) {
				t.Errorf("Got link %v, Want link %v", sl.head, want.head)
			}

			t.Logf("Got Link: %s, Want Link: %s", sl, want)
		})

	}
}

func TestSingleLink_RemoveNext(t *testing.T) {
	type fields struct {
		size int
		head *SingleLinkNode
		tail *SingleLinkNode
	}
	type args struct {
		node *SingleLinkNode
	}
	sl := makeSingleLinkWithArray([]int{1, 2, 3, 4, 5})
	tests := []struct {
		name    string
		fields  *SingleLink
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Remove a node.",
			fields: sl,
			args: args{
				node: &SingleLinkNode{
					Val: 3,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := tt.fields
			t.Logf("Old Link: %s\n", sl.head)
			got, _ := sl.RemoveNext(tt.args.node)
			if got != tt.want {
				t.Errorf("SingleLink.RemoveNext() = %v, want %v", got, tt.want)
			}
			t.Logf("New link: %s\n", sl.head)
		})
	}
}
