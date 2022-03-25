package creational

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	var s1, s2 Singleton
	s1 = GetInstance()
	s2 = GetInstance()
	s1.SetX(1)
	s2.SetY(2)
	if s1.GetX() != s2.GetX() {
		t.Errorf("Instance Var Not Same: %d, %d", s1.GetX(), s2.GetX())
	}
	if s1.GetY() != s2.GetY() {
		t.Errorf("Instance Var Not Same: %d, %d", s1.GetY(), s2.GetY())
	}
}
