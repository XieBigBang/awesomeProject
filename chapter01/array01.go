package chapter01

import "fmt"

func Array01() {
	//array := [5]*int{0: new(int), 1:new(int)}
	array := [5]int{1, 22, 31, 11, 4}

	array[0] = 10
	array[1] = 20

	for i, v := range array {
		fmt.Printf("array index:%d  value:%d\n", i, v)
	}
}
