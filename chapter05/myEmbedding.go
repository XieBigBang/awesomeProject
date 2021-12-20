package chapter05

import "fmt"

type mover interface {
	move()
}

type animal struct {
	name string
	size string
}

func (a *animal) SayName() {
	fmt.Printf("animal name:%v, size:%v\n", a.name, a.size)
}

func (a *animal) move() {
	fmt.Printf("animal moveeeeeeeeeee\n")
}

type bird struct {
	animal
	Color string
}

func NewBird(name string) *bird {
	return &bird{animal{name, "small"}, "yellow"}
}

func (b bird) SayColor() {
	fmt.Printf("bird name:%v  color:%v\n", b.name, b.Color)
}

func Test0509() {
	b := bird{animal{"little bird", "small"}, "colorful"}
	b.SayName()
	b.animal.SayName()

	goMove(&b)
}

func goMove(m mover) {
	m.move()
}
