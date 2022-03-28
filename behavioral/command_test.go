package behavioral

import "testing"

func TestCommand(t *testing.T) {
	a := aircraft{}
	p := phone{command: takeShot{a}}
	ra := radioTransmitter{command: fly{a}}
	p.Send()
	ra.Send()
}
