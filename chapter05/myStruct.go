package chapter05

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

type Duration int

func Test0501() {
	var user01 user
	fmt.Println(user01)
	fmt.Println(user01.name)
	fmt.Println(user01.ext)

	user02 := user{
		name:       "lisa",
		email:      "xxx@xx.com",
		ext:        2,
		privileged: true,
	}
	_ = user02
	fmt.Println(user02)

	tom := admin{
		person: user{
			name:       "tom",
			email:      "xxxxx",
			ext:        1,
			privileged: true,
		},
		level: "high",
	}
	fmt.Println(tom)

	var dd Duration
	dd = Duration(12)
	fmt.Println(dd)

	var ddd Duration
	//ddd = int(21)
	_ = ddd
}
