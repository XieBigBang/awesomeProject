package chapter05

import "fmt"

type sounder interface {
	sound()
}

type bee struct {
	name string
	size string
}

func (b bee) sound() {
	fmt.Printf("bee %v buzzzzzzzzzzzz\n", b.name)
}

type dog struct {
	name  string
	color string
}

func (d dog) sound() {
	fmt.Printf("dog %v woof woof\n", d.name)
}

func Test0508() {
	b := bee{"little bee", "small"}
	animalMakeSound(b)

	d := dog{"yellow dog", "yellow"}
	animalMakeSound(d)
}

func animalMakeSound(s sounder) {
	s.sound()
}
