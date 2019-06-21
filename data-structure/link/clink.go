package link

type CircularSingleLink struct {
	size int
	head *CircularSingleLinkNode
}

type CircularSingleLinkNode struct {
	Val  int
	Next *CircularSingleLinkNode
}

func NewCircularSingleLink() *CircularSingleLink {
	return new(CircularSingleLink)
}

func (csl *CircularSingleLink) InsertNext(node *CircularSingleLinkNode, data int) error {
	newNode := &CircularSingleLinkNode{
		Val: data,
	}

	if csl.size == 0 {
		// 第一个元素
		if node != nil {
			return ERROR_DOES_NOT_EXIST
		}
		newNode.Next = newNode
		csl.head = newNode
	} else {
		newNode.Next = node.Next
		node.Next = newNode
	}
	csl.size++
	return nil
}

func (csl *CircularSingleLink) RemoveNext(node *CircularSingleLinkNode) (int, error) {
	if csl.size == 0 {
		return 0, ERROR_EMPTY_LINK
	}

	var ret int
	if node.Next == node {
		// 最后一个节点
		csl.head = nil
		ret = node.Val
	} else {
		old := node.Next
		ret = node.Next.Val
		node.Next = node.Next.Next
		if old == csl.head {
			csl.head = node
		}
	}
	csl.size--
	return ret, nil
}
func (csl *CircularSingleLink) Size() int {
	return csl.size
}

func (csl *CircularSingleLink) Head() *CircularSingleLinkNode {
	return csl.head
}

func (csl *CircularSingleLink) isHead(node *CircularSingleLinkNode) bool {
	return csl.head == node
}
