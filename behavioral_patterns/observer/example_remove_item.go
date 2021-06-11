package main

import "fmt"

func removeItemfromSlice(list []int, num int) []int {
	fmt.Println(list)
	fmt.Println("remove int i from the list", num)

	for i, n := range list {
		if n == num {
			//swap n with last elements of the list
			list[len(list)-1], list[i] = list[i], list[len(list)-1]
			return list[:len(list)-1]
		}
	}
	return list
}

func run() {

	list := []int{10, 2, 3}

	Newlist := removeItemfromSlice(list, 10)
	fmt.Println(Newlist)

	//	for _, l := range Newlist {
	//		fmt.Println(l)
	//	}

}
