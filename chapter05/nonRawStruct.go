package chapter05

import "fmt"

type tax struct {
	num       string
	positionX int16
	positionY int16
}

type Tax struct {
	*tax
}

func NewTax(num string) (t *Tax, err error) {
	t2 := &Tax{
		tax: &tax{
			"", 0, 0,
		},
	}
	t2.num = num
	t2.positionX = 0
	t2.positionY = 0
	return t2, nil
}

func (t *Tax) Move(x int16, y int16) {
	t.positionY = y
	t.positionX = x
}

func Test0503() {
	t, _ := NewTax("YB6666")
	fmt.Printf("tax num=%v, pX=%v, pY=%v \n", t.num, t.positionX, t.positionY)

	t.Move(21, 223)
	fmt.Printf("tax num=%v, new pX=%v, new pY=%v \n", t.num, t.positionX, t.positionY)

}
