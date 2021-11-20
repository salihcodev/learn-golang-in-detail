package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	fmt.Println(arr[:2])  // [a b]
	fmt.Println(arr[3:])  // [d e f g h]
	fmt.Println(arr[3:6]) // [d e f]

	test := [...]int{1, 2}
	fmt.Println(test)

}
