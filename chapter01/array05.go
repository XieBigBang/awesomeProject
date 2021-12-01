package chapter01

import "fmt"

func Array05() {
	var arr01 [4][2]int
	arr01 = [4][2]int{{1, 2}, {10, 20}, {11, 22}, {2, 4}}
	for i, v := range arr01 {
		fmt.Printf("%d  %d\n", i, v)
	}

	arr02 := [4][2]int{1: {10, 20}, 3: {22, 11}}
	for i, v := range arr02 {
		fmt.Printf("%d  %d\n", i, v)
	}

	arr02[0] = [2]int{23, 24}
	arr02[2][1] = 44
	for i, v := range arr02 {
		for j, k := range v {
			fmt.Printf("[%d, %d]  %d\n", i, j, k)
		}
	}
}
