package link

import (
	"fmt"
)

func makeSingleLinkWithArray(data []int) *SingleLink {
	dummy := new(SingleLinkNode)
	p := dummy
	for i := range data {
		newNode := &SingleLinkNode{
			Val: data[i],
		}
		p.Next = newNode
		p = p.Next
	}
	return &SingleLink{
		size: len(data),
		head: dummy.Next,
		tail: p,
	}
}

func makeDoubleLinkWithArray(data []int) *DoubleLink {
	dummy := new(DoubleLinkNode)
	p := dummy
	for i := range data {
		newNode := &DoubleLinkNode{
			Val:  data[i],
			Prev: p,
		}
		p.Next = newNode
		p = p.Next
	}
	dummy = dummy.Next
	dummy.Prev = nil
	return &DoubleLink{
		size: len(data),
		head: dummy,
		tail: p,
	}
}

func makeCircularSingleLinkWithArray(data []int) *CircularSingleLink {
	dummy := new(CircularSingleLinkNode)
	p := dummy
	for i := range data {
		newNode := &CircularSingleLinkNode{
			Val: data[i],
		}
		p.Next = newNode
		p = p.Next
	}
	p.Next = dummy.Next
	return &CircularSingleLink{
		size: len(data),
		head: dummy.Next,
	}
}

func (sl *SingleLink) String() string {
	return fmt.Sprintf("{\n"+
		"Size: %d,\n"+
		"Head: %s,\n"+
		"}\n", sl.size, sl.head)
}

func (sln *SingleLinkNode) String() string {
	var result string
	var temp []interface{}
	dummy := sln
	for dummy != nil {
		result = result + "%d -> "
		temp = append(temp, dummy.Val)
		dummy = dummy.Next
	}
	result = result + "nil"

	return fmt.Sprintf(result, temp...)
}

func (dl *DoubleLink) String() string {
	return fmt.Sprintf("{\n"+
		"Size: %d,\n"+
		"Head: %s,\n"+
		"}\n", dl.size, dl.head)
}

func (dln *DoubleLinkNode) String() string {
	var result string
	result = "nil <=> "
	var temp []interface{}
	dummy := dln
	for dummy != nil {
		result = result + "%d <=> "
		temp = append(temp, dummy.Val)
		dummy = dummy.Next
	}
	result = result + "nil"

	return fmt.Sprintf(result, temp...)
}

func (csl *CircularSingleLink) String() string {
	return fmt.Sprintf("{\n"+
		"Size: %d,\n"+
		"Head: %s,\n"+
		"}\n", csl.size, csl.head)
}

func (csln *CircularSingleLinkNode) String() string {
	var result string
	var temp []interface{}
	dummy := csln
	count := 0
	for {
		if count > 0 && dummy == csln {
			break
		}
		if count == 0 {
			count++
		}

		result = result + "%d -> "
		temp = append(temp, dummy.Val)
		dummy = dummy.Next
	}
	result = result + "head"

	return fmt.Sprintf(result, temp...)
}
