package creational

import "testing"

func TestPrototype(t *testing.T) {
	var l1 *linkedList
	for i := 0; i < 3; i++ {
		l1 = &linkedList{i, l1}
	}
	l2 := l1.Clone()
	l1.next.val = 200
	if l2.Next().GetVal() == l1.Next().GetVal() {
		t.Fatal("clone did not perform deep copy")
	}
}
