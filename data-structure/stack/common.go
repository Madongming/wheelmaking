package stack

import (
	"fmt"
)

func (s *Stack) String() string {
	return fmt.Sprintf("{\n"+
		"Size: %d,\n"+
		"Head: %s,\n"+
		"}\n", s.size, s.head)
}

func (sn *LinkNode) String() string {
	var result string
	var temp []interface{}
	dummy := sn
	for dummy != nil {
		result = result + "%d <- "
		temp = append(temp, dummy.Val)
		dummy = dummy.Next
	}
	result = result + "stack botton"

	return fmt.Sprintf(result, temp...)
}
