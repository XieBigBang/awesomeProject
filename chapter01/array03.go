package chapter01

import "fmt"

func Array03() {
	var arr01 [5]string

	arr02 := [5]string{"a", "b", "c", "d", "e"}

	arr01 = arr02

	for v, i := range arr01 {
		fmt.Printf("%d  %s\n", v, i)
	}
}
