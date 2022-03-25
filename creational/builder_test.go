package creational

import "testing"

func TestBuilder(t *testing.T) {
	b := GetBuilder()
	b.AddLens("Zoom Lens", "70-200mm", 2.8)
	b.AddBody("Focal Plane", "CMOS", "Metal")
	c := b.GetProduct()
	if c == nil {
		t.Fatal("Expecting Product, got nil")
	} else {
		c.Briefing()
	}
}
