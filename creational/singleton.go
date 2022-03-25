package creational

import "sync"

var _singleton_once sync.Once
var singletonInstance Singleton

type singleton struct {
	x int
	y int
}

type Singleton interface {
	GetX() int
	GetY() int
	SetX(newX int)
	SetY(newY int)
}

func GetInstance() Singleton {
	if singletonInstance == nil {
		_singleton_once.Do(
			func() {
				singletonInstance = &singleton{0, 0}
			})
	}
	return singletonInstance
}

func (p *singleton) GetX() int {
	return p.x
}

func (p *singleton) GetY() int {
	return p.y
}

func (p *singleton) SetX(newX int) {
	p.x = newX
}

func (p *singleton) SetY(newY int) {
	p.y = newY
}
