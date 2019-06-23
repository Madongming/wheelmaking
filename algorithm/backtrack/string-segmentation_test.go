package backtrack

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "partison abcd",
			args: args{
				str: "abcd",
			},
			want: [][]string{[]string{"a", "b", "c", "d"},
				[]string{"a", "b", "cd"},
				[]string{"a", "bc", "d"},
				[]string{"a", "bcd"},
				[]string{"ab", "c", "d"},
				[]string{"ab", "cd"},
				[]string{"abc", "d"},
				[]string{"abcd"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.args.str)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partision() = %v, want %v", got, tt.want)
			}
			t.Logf("The string \"%s\" 's result of the segmentation is\n %v", tt.args.str, got)
		})
	}
}
