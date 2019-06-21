package link

type SingleLinkNode struct {
	Val  int
	Next *SingleLinkNode
}

type SingleLink struct {
	size int

	head *SingleLinkNode
	tail *SingleLinkNode
}

func NewSingleLink() *SingleLink {
	return new(SingleLink)
}

func (sl *SingleLink) InsertNext(node *SingleLinkNode, data int) error {
	// 链表为空，选择点不为空，一定为异常
	if sl.size == 0 && node != nil {
		return ERROR_DOES_NOT_EXIST
	}

	// 加入的点
	newNode := &SingleLinkNode{
		Val: data,
	}

	if sl.size == 0 {

		// 链表和node均为空，说明是要向一个空链表中插入数据
		sl.head = newNode
		sl.tail = newNode
		sl.size++
		return nil
	} else if node == nil {
		// 将数据插入非空链表的头节点前
		newNode.Next = sl.head
		sl.head = newNode
		sl.size++
		return nil
	}

	dummy := sl.head
	for dummy != nil {
		if dummy.Val == node.Val {
			newNode.Next = dummy.Next
			dummy.Next = newNode
			sl.size++

			if newNode.Next == nil {
				sl.tail = newNode
			}
			return nil
		}
		dummy = dummy.Next
	}
	return ERROR_DOES_NOT_EXIST
}

func (sl *SingleLink) RemoveNext(node *SingleLinkNode) (int, error) {
	if sl.size == 0 {
		return 0, ERROR_EMPTY_LINK
	}

	if node == nil {
		// 删除头节点
		ret := sl.head.Val
		sl.head = sl.head.Next
		sl.size--
		return ret, nil
	}

	dummy := sl.head
	for dummy != nil {
		if dummy.Val == node.Val {
			if dummy.Next == nil {
				return 0, ERROR_DOES_NOT_EXIST
			}
			ret := dummy.Next.Val
			dummy.Next = dummy.Next.Next
			if dummy.Next == nil {
				sl.tail = dummy
			}
			return ret, nil
		}
		dummy = dummy.Next
	}
	return 0, ERROR_DOES_NOT_EXIST
}

func (sl *SingleLink) Size() int {
	return sl.size
}

func (sl *SingleLink) Head() *SingleLinkNode {
	return sl.head
}

func (sl *SingleLink) Tail() *SingleLinkNode {
	return sl.tail
}

func (sl *SingleLink) isHead(node *SingleLinkNode) bool {
	return sl.head == node
}

func (sl *SingleLink) isTail(node *SingleLinkNode) bool {
	return sl.tail == node
}
