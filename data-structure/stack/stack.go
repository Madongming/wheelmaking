package stack

type Stack struct {
	head *LinkNode
	tail *LinkNode
	size int
}

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(data int) error {
	newNode := &LinkNode{
		Val: data,
	}

	if s.size == 0 {
		// 第一个元素
		s.head = newNode
		s.tail = newNode
	} else {
		newNode.Next = s.head
		s.head = newNode
	}
	s.size++
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return 0, ERROR_STACK_IS_EMPTY
	}

	ret := s.head.Val
	s.head = s.head.Next
	s.size--

	return ret, nil
}

func (s *Stack) Peek() (int, error) {
	if s.size == 0 {
		return 0, ERROR_STACK_IS_EMPTY
	}

	return s.head.Val, nil
}
