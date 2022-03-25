package creational

import "testing"

func TestFactory(t *testing.T) {
	d, _ := Factory("display")
	p, _ := Factory("phoneScreen")
	_, err := Factory("Whatever Product")
	d.Introduce()
	// Output:
	// I am a 27 inch Display!
	p.Introduce()
	// Output:
	// I am a 5 inch PhoneScreen!
	if err == nil {
		t.Fatal("Got nil error for invalid type")
	}
}
