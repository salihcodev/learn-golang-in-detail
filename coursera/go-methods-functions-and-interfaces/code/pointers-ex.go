package main

import "fmt"


//# arrays need to be passed with the pointer

//func test(x *[4]int) {
//	(*x)[0] += (*x)[0]
//}
//
//func main() {
//	a := [4]int{1, 2, 3, 5}
//
//	fmt.Println(a)
//
//	test(&a)
//
//	fmt.Println(a)
//}


//# on the other hand the slices are passed by reference by default. AKA 'no need to pass the pointer'
func test(x []int) {
	(x)[0] += (x)[0]
}

func main() {
	a := []int{1, 2, 3, 5}

	fmt.Println(a)

	test(a)

	fmt.Println(a)
}
