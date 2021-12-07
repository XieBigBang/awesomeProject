package chapter01

import (
	"fmt"
)

func Map01() {
	dict := make(map[string]int)
	fmt.Printf("map len %v\n", len(dict))
	fmt.Printf("================\n")

	dict01 := map[string]int{"a": 1, "b": 22, "c": 32, "d": 2, "e": 32}
	for k, v := range dict01 {
		fmt.Printf("k:%v v:%v\n", k, v)
	}

	//dict01 := map[[]string]int{}
	//dict01 := map[int][]string{}

	dict01["a"] = 100
	fmt.Printf("dict[%v]=%v\n", "a", dict01["a"])
	fmt.Printf("dict[%v]=%v\n", "e", dict01["e"])
	fmt.Printf("================\n")

	//var colors map[string]string
	//colors["aaa"] = "asdf"   // assignment to entry in nil map

	val, exists := dict01["dd"]
	if !exists {
		fmt.Printf("key dd is not exists\n")
	} else {
		fmt.Printf("v %v\n", val)
	}

	delete(dict01, "a")
	val, exists = dict01["a"]
	if !exists {
		fmt.Printf("a has delete\n")
	}

	dict01["f"] = 232
	for k, v := range dict01 {
		fmt.Printf("k %v  v %v\n", k, v)
	}
	fmt.Printf("================\n")

	removeEle(dict01, "f")
	for k, v := range dict01 {
		fmt.Printf("k %v  v %v\n", k, v)
	}

}

func removeEle(amap map[string]int, key string) {
	delete(amap, key)
}
