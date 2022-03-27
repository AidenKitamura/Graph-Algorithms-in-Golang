package creational

import (
	"fmt"
	"sync"
)

var once sync.Once

var P *pool

type object struct {
	name string
}

func (o object) GetName() string {
	return o.name
}

type Object interface {
	GetName() string
}

type pool struct {
	*sync.Mutex
	idle   []object
	active []object
}

func (p *pool) Acquire(name string) (Object, error) {
	p.Lock()
	defer p.Unlock()
	for i := 0; i < len(p.idle); i++ {
		if p.idle[i].name == name {
			res := p.idle[i]
			p.idle[i] = p.idle[len(p.idle)-1]
			p.idle = p.idle[:len(p.idle)-1]
			p.active = append(p.active, res)
			return res, nil
		}
	}
	return nil, fmt.Errorf("cannot find object with name %s", name)
}

func (p *pool) Free(name string) error {
	p.Lock()
	defer p.Unlock()
	for i := 0; i < len(p.active); i++ {
		if p.active[i].name == name {
			res := p.active[i]
			p.active[i] = p.active[len(p.active)-1]
			p.active = p.active[:len(p.active)-1]
			p.idle = append(p.idle, res)
			return nil
		}
	}
	return fmt.Errorf("cannot free object with name %s", name)
}

type Pool interface {
	Acquire(name string) (Object, error)
	Free(name string) error
}

func InitPool(cap int) Pool {
	once.Do(func() {
		P = &pool{}
		P.Mutex = new(sync.Mutex)
		P.active = make([]object, 0)
		P.idle = make([]object, 0)
		for i := 0; i < cap; i++ {
			P.idle = append(P.idle, object{fmt.Sprintf("%d", i)})
		}
	})
	return P
}
