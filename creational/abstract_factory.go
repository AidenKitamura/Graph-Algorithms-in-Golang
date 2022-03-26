package creational

import "fmt"

type AbstractFactory interface {
	ProducePhone() Phone
}

type Phone struct {
	brand string
}

type samFactory struct {
}

type applFactory struct {
}

func (s samFactory) ProducePhone() Phone {
	return Phone{"Sam"}
}

func (a applFactory) ProducePhone() Phone {
	return Phone{"Appl"}
}

func GetAbstractFactory(brand string) (AbstractFactory, error) {
	switch brand {
	case "Sam":
		return samFactory{}, nil
	case "Appl":
		return applFactory{}, nil
	}
	return nil, fmt.Errorf("invalid brand name %s", brand)
}
