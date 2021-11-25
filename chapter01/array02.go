package chapter01

import "fmt"

func Array02() {
	array := [5]*int{0: new(int), 1: new(int)}

	*array[0] = 10
	*array[1] = 20

	for i, v := range array {
		if v != nil {
			fmt.Printf("array index:%d  value:%d\n", i, *v)
		} else {
			fmt.Printf("array index:%d  value:nil \n", i)
		}
	}
}
