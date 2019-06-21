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
	if dl.size != 0 && node == nil {
		// 如果双向链表不为空，向空节点后加入与向头节点前插入相同，故这里禁止
		return ERROR_DOES_NOT_EXIST
	}

	newNode := &DoubleLinkNode{
		Val: data,
	}

	if dl.size == 0 {
		// 第一个元素
		dl.head = newNode
		dl.tail = newNode
	} else {
		newNode.Prev = node
		newNode.Next = node.Next
		if newNode.Next == nil {
			dl.tail = newNode
		} else {
			node.Next.Prev = newNode
		}
		node.Next = newNode
	}
	dl.size++
	return nil
}

func (dl *DoubleLink) InsertPrev(node *DoubleLinkNode, data int) error {
	if dl.size != 0 && node == nil {
		return ERROR_DOES_NOT_EXIST
	}

	newNode := &DoubleLinkNode{
		Val: data,
	}

	if dl.size == 0 {
		// 第一个元素
		dl.head = newNode
		dl.tail = newNode
	} else {
		newNode.Prev = node.Prev
		newNode.Next = node
		if newNode.Prev == nil {
			dl.head = newNode
		} else {
			node.Prev.Next = newNode
		}
		node.Prev = newNode
	}
	dl.size++
	return nil
}

func (dl *DoubleLink) Remove(node *DoubleLinkNode) (int, error) {
	if node == nil || dl.size == 0 {
		return 0, ERROR_EMPTY_LINK
	}

	ret := node.Val

	if dl.size == 1 {
		// 只有一个元素
		dl.head = nil
		dl.tail = nil
	} else if node == dl.head {
		// 元素在首位
		dl.head = node.Next
		node.Next.Prev = nil
	} else if node == dl.tail {
		// 元素在尾部
		dl.tail = node.Prev
		node.Prev.Next = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
	dl.size--
	return ret, nil
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
