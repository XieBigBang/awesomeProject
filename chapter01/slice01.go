package chapter01

import "fmt"

func Slice01() {
	sl02 := make([]int, 3)
	for i, v := range sl02 {
		fmt.Printf("%d  %d\n", i, v)
	}
	fmt.Printf("=============\n")

	//sl01 := make([]*string,3)
	sl01 := []*string{0: new(string), 2: new(string)}

	*sl01[0] = "aa"
	//sl01[1] = new(string)
	//*sl01[1] = "aa"
	*sl01[2] = "aa"

	for i, v := range sl01 {
		if v != nil {
			fmt.Printf("%d  %s\n", i, *v)
		}
	}
	fmt.Printf("=============\n")

	sl03 := []int{1, 2, 3, 3}
	sl03[0] = 111
	for i, v := range sl03 {
		fmt.Printf("%d  %d\n", i, v)
	}
	fmt.Printf("%d\n", sl03)
	fmt.Printf("=============\n")

	var sl04 []int
	fmt.Printf("sl04 len %d cap %d\n", len(sl04), cap(sl04))
	fmt.Printf("sl04  %d\n", sl04)
	fmt.Printf("=============\n")

	sl05 := make([]int, 0)
	fmt.Printf("%d\n", sl05)
	fmt.Printf("sl05 len %d  cap %d\n", len(sl05), cap(sl05))

	fmt.Printf("=============\n")

	sl06 := []int{1, 2, 3, 4, 5, 6}
	sl07 := sl06[1:3]
	fmt.Printf("sl06 len %d cap %d\n", len(sl06), cap(sl06))
	// cap(slice) = k, newslice := slice[i:j]
	// len(newslice) = j - i
	// cap(newslice) = k - i
	fmt.Printf("sl07 len %d cap %d\n", len(sl07), cap(sl07)) // sl07 len 2 cap 5
	for i, v := range sl07 {
		fmt.Printf("%d  %d\n", i, v)
	}
	fmt.Printf("--------------\n")
	sl07[0] = 888
	for i, v := range sl06 {
		fmt.Printf("%d  %d\n", i, v)
	}

	//sl07[2] = 33  // runtime error: index out of range [2] with length 2

	fmt.Printf("=============\n")
	sl08 := append(sl07, 555)
	fmt.Printf("sl08[2] %d\n", sl08[2])
	fmt.Printf("sl07 len %d cap %d\n", len(sl07), cap(sl07))
	fmt.Printf("sl08 len %d cap %d\n", len(sl08), cap(sl08))

	fmt.Printf("=============\n")
	sl09 := append(sl07, 666, 666, 666)
	fmt.Printf("sl08[2] %d\n", sl08[2])
	fmt.Printf("sl09[2] %d\n", sl09[2])
	fmt.Printf("sl07 len %d cap %d\n", len(sl07), cap(sl07))
	fmt.Printf("sl09 len %d cap %d\n", len(sl09), cap(sl09))

	fmt.Printf("=============\n")
	sl10 := append(sl07, 777, 777, 777, 777)
	fmt.Printf("sl08[2] %d\n", sl08[2]) // sl08[2] 666
	fmt.Printf("sl10[2] %d\n", sl10[2]) // sl10[2] 777
	fmt.Printf("sl10[0] %d\n", sl10[0]) // sl10[0] 888
	fmt.Printf("sl07 len %d cap %d\n", len(sl07), cap(sl07))
	fmt.Printf("sl10 len %d cap %d\n", len(sl10), cap(sl10)) // sl10 len 6 cap 10

}
