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

// node为链表中的一个节点，在其之后插入数据，如果node为空表示在头节点前插入
func (sl *SingleLink) InsertNext(node *SingleLinkNode, data int) error {
	// 加入的点
	newNode := &SingleLinkNode{
		Val: data,
	}

	if node == nil {
		// node为空，说明是要向头节点之前插入数据
		sl.head = newNode
		if sl.size == 0 {
			sl.tail = newNode
		}
	} else {
		newNode.Next = node.Next
		node.Next = newNode
		if node.Next == nil {
			sl.tail = node
		}
	}
	sl.size++
	return nil
}

func (sl *SingleLink) RemoveNext(node *SingleLinkNode) (int, error) {
	if sl.size == 0 {
		return 0, ERROR_EMPTY_LINK
	}

	var ret int
	if node == nil {
		// 删除头节点
		if sl.size == 1 {
			sl.head = nil
			sl.tail = nil
		}
		ret = sl.head.Val
		sl.head = sl.head.Next
	} else {
		if node.Next == nil {
			return 0, ERROR_DOES_NOT_EXIST
		}
		ret = node.Next.Val
		node.Next = node.Next.Next
		if node.Next == nil {
			sl.tail = node
		}
	}
	sl.size--
	return ret, nil
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
