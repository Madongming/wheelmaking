package link

type DoubleLinkNode struct {
	Val  int
	Prev *DoubleLinkNode
	Next *DoubleLinkNode
}

type DoubleLink struct {
	size int

	head *DoubleLinkNode
	tail *DoubleLinkNode
}

func NewDoubleLink() *DoubleLink {
	return new(DoubleLink)
}

func (dl *DoubleLink) InsertNext(node *DoubleLinkNode, data int) error {
	if dl.size == 0 && node != nil {
		// 向指定位置的空链表中插入中数据，一定异常
		return ERROR_DOES_NOT_EXIST
	}

	newNode := &DoubleLinkNode{
		Val: data,
	}

	if dl.size == 0 {
		// 向未指定位置的空列表中插入数据，第一个元素
		dl.head = newNode
		dl.tail = newNode
		dl.size++
		return nil
	} else if node == nil {
		// 向未指定位置的空列表中插入数据，插在头节点之前
		newNode.Next = dl.head
		dl.head.Prev = newNode
		dl.head = newNode
		dl.size++
		return nil
	}

	dummy := dl.head
	for dummy != nil {
		if dummy.Val == node.Val {
			newNode.Next = dummy.Next
			newNode.Prev = dummy
			if dummy.Next == nil {
				dl.tail = newNode
			} else {
				dummy.Next.Prev = newNode
			}
			dummy.Next = newNode

			dl.size++
			return nil
		}
		dummy = dummy.Next
	}
	return ERROR_DOES_NOT_EXIST
}

func (dl *DoubleLink) InsertPrev(node *DoubleLinkNode, data int) error {
	if dl.size == 0 && node != nil {
		return ERROR_DOES_NOT_EXIST
	}

	newNode := &DoubleLinkNode{
		Val: data,
	}

	if dl.size == 0 {
		dl.head = newNode
		dl.tail = newNode
		dl.size++
		return nil
	} else if node == nil {
		// 插在尾节点之后
		newNode.Prev = dl.tail
		dl.tail.Next = newNode
		dl.tail = newNode
		dl.size++
		return nil
	}

	dummy := dl.head
	for dummy != nil {
		if dummy.Val == node.Val {
			newNode.Next = dummy
			newNode.Prev = dummy.Prev
			if dummy.Prev == nil {
				dl.head = newNode
			} else {
				dummy.Prev.Next = newNode
			}
			dummy.Prev = newNode

			dl.size++
			return nil
		}
		dummy = dummy.Next
	}
	return ERROR_DOES_NOT_EXIST
}

func (dl *DoubleLink) Remove(node *DoubleLinkNode) (int, error) {
	if dl.size == 0 {
		return 0, ERROR_EMPTY_LINK
	}

	dummy := dl.head
	for dummy != nil {
		if dummy.Val == node.Val {
			dl.size--
			if dummy.Next == nil {
				// 元素在首位
				dl.tail = dummy.Prev
				dummy.Prev.Next = nil
				return dummy.Val, nil
			} else if dummy.Prev == nil {
				// 元素在尾部
				dl.head = dummy.Next
				dummy.Next.Prev = nil
				return dummy.Val, nil
			}

			dummy.Next.Prev = dummy.Prev
			dummy.Prev.Next = dummy.Next
			return dummy.Val, nil
		}
		dummy = dummy.Next
	}
	return 0, ERROR_DOES_NOT_EXIST
}

func (dl *DoubleLink) Size() int {
	return dl.size
}

func (dl *DoubleLink) Head() *DoubleLinkNode {
	return dl.head
}

func (dl *DoubleLink) Tail() *DoubleLinkNode {
	return dl.tail
}

func (dl *DoubleLink) isHead(node *DoubleLinkNode) bool {
	return dl.head == node
}

func (dl *DoubleLink) isTail(node *DoubleLinkNode) bool {
	return dl.tail == node
}
