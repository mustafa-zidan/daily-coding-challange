package structures

type LinkedListNode struct {
	Val  int
	Next *LinkedListNode
}

func (n *LinkedListNode) New(items []int) *LinkedListNode {
	root := &LinkedListNode{}
	return root
}

func (l *LinkedListNode) Add(item int) {
	next := l.Next
	if next != nil {
		for next.Next != nil {
			next = next.Next
		}
	}
	next = &LinkedListNode{item, nil}
	l.Next = next
}

func (l *LinkedListNode) Remove(i int) {
	node := l.Get(i - 1)
	if node.Next != nil {
		node.Next.Next = node.Next
	}
}

func (l *LinkedListNode) Get(i int) *LinkedListNode {
	item := l.Next
	for count := 0; count <= i; count++ {
		item = item.Next
	}
	return item
}

func (l *LinkedListNode) Len() int {
	count := 1
	next := l.Next
	for next != nil {
		next = next.Next
		count++
	}
	return count
}
