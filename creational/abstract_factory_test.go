package creational

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	af, err := GetAbstractFactory("Appl")
	if err != nil {
		t.Fatal("Initialization failed. What could be the issue?")
	} else {
		fmt.Println("Got a new Phone: ", af.ProducePhone().brand)
	}
	_, err = GetAbstractFactory("Some Unknown Brand")
	if err == nil {
		t.Fatal("Expecting Error, but got nil")
	} else {
		fmt.Println("Got expected error message: ", err)
	}
}
