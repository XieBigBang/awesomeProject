package chapter05

type human struct {
	Name  string
	Email string
}

type Teacher struct {
	human
	Level string
}
