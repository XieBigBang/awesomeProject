package chapter04

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
	fmt.Printf("%v\n", sl03)
	fmt.Printf("sl03 len %v  cap %v\n", len(sl03), cap(sl03))
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

	fmt.Printf("=============\n")
	sl11 := []string{"a", "b", "c", "d", "e"}
	// newslice := slice[i:j:k]
	// len(newslice) = j-i
	// cap(newslice) = k-i
	sl12 := sl11[1:2:3]
	// sl12 := sl11[1:2:6]  // runtime error: slice bounds out of range [::6] with capacity 5
	sl12_ := append(sl12, "bbb") // appends elements to the end of a slice
	sl12_[0] = "kkkk"

	fmt.Printf("sl11 len %d  cap %d\n", len(sl11), cap(sl11)) // sl11 len 5  cap 5
	fmt.Printf("sl12 len %d  cap %d\n", len(sl12), cap(sl12)) // sl12 len 1  cap 2
	fmt.Printf("sl11 [] = %v\n", sl11)
	fmt.Printf("sl12 [] = %v\n", sl12)
	fmt.Printf("append sl12  = %v\n", sl12_)
	fmt.Printf("sl11 [] = %v\n", sl11)
	fmt.Printf("sl12 [] = %v\n", sl12)

	fmt.Printf("=============\n")
	sl13 := []string{"a", "b", "c", "d", "d", "e"}
	sl14 := sl13[3:4:4]
	sl15 := append(sl14, "aaaa") // a new underlying array
	sl15[0] = "bbbb"
	fmt.Printf("sl13 = %v\n", sl13)
	fmt.Printf("sl14 = %v\n", sl14)
	fmt.Printf("sl15 = %v\n", sl15)

	fmt.Printf("=============\n")
	sl16 := []int{1, 2, 3, 4, 5, 6}
	for i, v := range sl16 {
		fmt.Printf("i=%d v=%d &v=%v &sl16[%d]=%v\n", i, v, &v, i, &sl16[i])
	}
	/**
	i=0 v=1 &v=0xc00000a138 &sl16[0]=0xc00000c2d0
	i=1 v=2 &v=0xc00000a138 &sl16[1]=0xc00000c2d8
	i=2 v=3 &v=0xc00000a138 &sl16[2]=0xc00000c2e0
	i=3 v=4 &v=0xc00000a138 &sl16[3]=0xc00000c2e8
	i=4 v=5 &v=0xc00000a138 &sl16[4]=0xc00000c2f0
	i=5 v=6 &v=0xc00000a138 &sl16[5]=0xc00000c2f8

	v 是副本，迭代过程中每个元素的副本存在同一个地址里
	*/
}
