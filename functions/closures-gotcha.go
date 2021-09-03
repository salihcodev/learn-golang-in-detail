package main

import "fmt"

func main() {

	f := make([]func(), 5)

	for i := 0; i < 5; i++ {
		_i := i

		/*
			__NOTE#2
			to fix it we gonna make a new variable 'i', then passing it to closures
		*/

		f[i] = func() {
			fmt.Printf("%p, %v\n", &_i, _i)
		}

		/*
			__NOTE#1
			after this loop finishes, all closures refer to the same variable 'i' which is by
			the end of the loop its value became = 4, then all closures has 'i' referred to 4
			as a value.


			to fix it see __NOTE#2

		*/
	}

	for i := 0; i < 5; i++ {
		f[i]()

		// it reduces:
		// 5
		// 5
		// 5
		// 5
		// 5

		/*
			so to fix that we need to know what causes it in the first place
			see *__NOTE#1*
		*/

	}

}
