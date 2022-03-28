package behavioral

import "fmt"

// Aircraft as receiver
type aircraft struct {
}

func (a aircraft) Fly() {
	fmt.Println("Flying!")
}

func (a aircraft) TakePhoto() {
	fmt.Println("Taking Photos!")
}

type Receiver interface {
	Fly()
	TakePhoto()
}

// we can use radiotranmitter or
// phone to control it
type radioTransmitter struct {
	command Command
}

func (r radioTransmitter) Send() {
	r.command.Execute()
}

type phone struct {
	command Command
}

func (p phone) Send() {
	p.command.Execute()
}

type Invoker interface {
	Send()
}

// we can command it to fly
// or take photos
type fly struct {
	r Receiver
}

func (f fly) Execute() {
	f.r.Fly()
}

type takeShot struct {
	r Receiver
}

func (t takeShot) Execute() {
	t.r.TakePhoto()
}

type Command interface {
	Execute()
}
