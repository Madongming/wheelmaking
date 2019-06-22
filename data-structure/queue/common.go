package queue

import (
	"fmt"
)

func (q *Queue) String() string {
	return fmt.Sprintf("{\n"+
		"Size: %d,\n"+
		"Head: %s,\n"+
		"}\n", q.size, q.head)
}

func (qn *LinkNode) String() string {
	result := "Queue head <- "
	var temp []interface{}
	dummy := qn
	for dummy != nil {
		result = result + "%d <- "
		temp = append(temp, dummy.Val)
		dummy = dummy.Next
	}
	result = result + "queue tail"

	return fmt.Sprintf(result, temp...)
}
