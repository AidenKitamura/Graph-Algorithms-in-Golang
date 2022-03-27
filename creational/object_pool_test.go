package creational

import "testing"

func TestObjectPool(t *testing.T) {
	p := InitPool(10)
	ob, _ := p.Acquire("1")
	_, err := p.Acquire("1")
	if ob == nil {
		t.Fatal("expecting non nil object")
	}
	if err == nil {
		t.Fatal("expecting non nil error")
	}
	err = p.Free("1")
	if err != nil {
		t.Fatal("expecting nil error")
	}
	err = p.Free("1")
	if err == nil {
		t.Fatal("expecting non nil error")
	}
}
