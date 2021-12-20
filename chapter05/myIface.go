package chapter05

import "fmt"

type notifier interface {
	notify()
}

type book struct {
	name string
	page int32
}

func (b *book) notify() {
	fmt.Printf("book  name:%v, page:%v", b.name, b.page)
}

func Test0506() {
	b := book{"my live", 54}
	//sendNotify(b)  // cannot use b (type book) as type notifier in argument to sendNotify:
	// book does not implement notifier (notify method has pointer receiver)

	sendNotify(&b)
}

func sendNotify(n notifier) {
	n.notify()
}
