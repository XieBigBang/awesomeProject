package chapter01

import "fmt"

func Array04() {
	var arr01 [3]*string

	arr02 := [3]*string{new(string), new(string), new(string)}

	arr01 = arr02

	*arr02[0] = "a"
	*arr02[1] = "b"
	*arr02[2] = "c"

	for i, v := range arr01 {
		fmt.Printf("%d  %s\n", i, *v)
	}
}
