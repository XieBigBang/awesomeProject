package chapter05

import "fmt"

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %v", *d)
}

func Test0507() {
	//duration(32).pretty()  // cannot call pointer method on duration(32)
	// cannot take the address of duration(32)

}
