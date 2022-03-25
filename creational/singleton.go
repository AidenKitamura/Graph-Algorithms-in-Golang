package creational

import "sync"

var _singleton_once sync.Once
var SingletonInstance Singleton

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
	if SingletonInstance == nil {
		_singleton_once.Do(
			func() {
				SingletonInstance = &singleton{0, 0}
			})
	}
	return SingletonInstance
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
