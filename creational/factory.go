package creational

import "fmt"

type Product interface {
	Introduce()
}

type display struct {
	description string
}

type phoneScreen struct {
	description string
}

func (d display) Introduce() {
	fmt.Println(d.description, "Display!")
}

func (p phoneScreen) Introduce() {
	fmt.Println(p.description, "PhoneScreen!")
}

func Factory(t string) (p Product, err error) {
	switch t {
	case "display":
		p = display{"I am a 27 inch"}
	case "phoneScreen":
		p = phoneScreen{"I am a 5 inch"}
	default:
		err = fmt.Errorf("invalid Product Type: %s", t)
	}
	return p, err
}
