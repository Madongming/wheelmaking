package queue

type Queue struct {
	size int
	head *LinkNode
	tail *LinkNode
}

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func NewQueue() *Queue {
	return new(Queue)
}

func (q *Queue) EnQueue(data int) error {
	newNode := &LinkNode{
		Val: data,
	}

	if q.size == 0 {
		// 第一个元素
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.Next = newNode
		q.tail = newNode
	}
	q.size++
	return nil
}

func (q *Queue) DeQueue() (int, error) {
	if q.size == 0 {
		return 0, ERROR_QUEUE_IS_EMPTY
	}

	var ret int
	ret = q.head.Val
	q.head = q.head.Next
	if q.size == 1 {
		q.tail = nil
	}
	q.size--
	return ret, nil
}

func (q *Queue) Peek() (int, error) {
	if q.size == 0 {
		return 0, ERROR_QUEUE_IS_EMPTY
	}

	return q.head.Val, nil
}
