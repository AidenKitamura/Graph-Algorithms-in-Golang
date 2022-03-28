package behavioral

import (
	"fmt"
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	n404 := networkResponse{404, "You should not see this", "You should not see this"}
	nProper := networkResponse{204, "Text", "You can see this"}
	nUnknown := networkResponse{204, "Non-Proper Header", "You should not see this"}
	var p Parser = &statusParser{}
	var p2 Parser = &headerParser{}
	p2.SetNext(&bodyParser{})
	p.SetNext(p2)
	fmt.Println("------------testing 404----------")
	p.Parse(n404)
	fmt.Println("------------testing Proper----------")
	p.Parse(nProper)
	fmt.Println("------------testing Unknown----------")
	p.Parse(nUnknown)
}
