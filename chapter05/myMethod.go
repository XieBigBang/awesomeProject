package chapter05

import "fmt"

type car struct {
	name  string
	cType string
}

func (c car) notify() {
	fmt.Printf("car name=%v,type=%v \n", c.name, c.cType)
}

func (c *car) changeName(name string) {
	c.name = name
}

func (c car) changeType(ctype string) {
	c.cType = ctype
}

func (c car) changeType2(ctype string) car {
	c.cType = ctype
	return c
}

func Test0502() {
	cc := car{
		name:  "honda",
		cType: "1",
	}
	cc.notify()

	c2 := &car{"toyota", "2"}
	c2.notify()

	cc.changeName("honda aa")
	cc.notify()

	c2.changeName("toyota bb")
	c2.notify()

	cc.changeType("111")
	cc.notify()

	c2.changeType("222")
	c2.notify()

	c3 := cc.changeType2("111")
	c3.notify()

	c4 := c2.changeType2("222")
	c4.notify()
}

// 方法接收者，该用值接收者，还是指针接收者，  要看这个值的本质，
// 原始本质    ----> 不可变对象，这种类型的方法 一般都是值接收者，即调用该方法是传递值的副本  --->对应内置类型
// 非原始本质  -----> 可变对象，一般使用指针接收者     ---->对应引用类型
//      结构类型，既可以是原始的，也可以是非原始的
