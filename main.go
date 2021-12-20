package main

import (
	"awesomeProject/chapter05"
	"fmt"
)

func main() {
	//chapter04.Array01()
	//chapter04.Array02()
	//chapter04.Array03()
	//chapter04.Array04()
	//chapter04.Array05()
	//chapter04.Slice01()
	//chapter04.Map01()
	//chapter05.Test0501()
	//chapter05.Test0502()
	//chapter05.Test0503()
	//chapter05.Test0504()
	//chapter05.Test0505()
	//chapter05.Test0506()
	//chapter05.Test0507()
	//chapter05.Test0508()
	//chapter05.Test0509()
	//
	//b := chapter05.NewBird("cuckoo")
	//b.Color = "colorful"
	//b.SayColor()
	//b.SayName()

	//c := chapter05.alertCounter(10)
	c := chapter05.New(10)
	fmt.Printf("counter: %v\n", c)

	u := chapter05.User{
		Name: "kk",
		//email: "32@s",
	}
	fmt.Printf("user name:%v\n", u.Name)

	t := chapter05.Teacher{}
	t.Name = "Wang"
	t.Email = "wang@com"
	t.Level = "high"
	fmt.Println(t)

}
