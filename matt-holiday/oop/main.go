package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

func main() {
	var t IntSlice = []int{1, 5, 8}

	for i, v := range t {
		fmt.Println(i, ":", v)
	}

	fmt.Printf("%T %[1]v\n", t)
}

func (s IntSlice) String() string {
	var strs []string

	for _, v := range s {
		strs = append(strs, strconv.Itoa(v))
	}

	return "[" + strings.Join(strs, "--") + "]"
}
