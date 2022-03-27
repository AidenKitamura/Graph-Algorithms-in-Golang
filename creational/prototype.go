package creational

type linkedList struct {
	val  int
	next *linkedList
}

func (l *linkedList) GetVal() int {
	return l.val
}

func (l *linkedList) Next() LinkedList {
	return l.next
}

func (l *linkedList) Clone() LinkedList {
	startNode := &linkedList{}
	cur := startNode
	for ptr := l; ptr != nil; ptr = ptr.next {
		cur.val = ptr.val
		if ptr.next != nil {
			cur.next = &linkedList{}
			cur = cur.next
		} else {
			cur.next = nil
		}
	}
	return startNode
}

type LinkedList interface {
	GetVal() int
	Next() LinkedList
	Clone() LinkedList
}
