package chapter05

type alertCounter int

func New(v int) alertCounter {
	return alertCounter(v)
}

type User struct {
	Name  string
	email string
}
