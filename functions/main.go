package main

import "fmt"

func main() {

	c, err := doSomThing(5, 6)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("the result is: %v\n", c)

	// defering functions calls:
	defer dd()
	/*
		dd function gonna be called only when the hosted -AKA main- function gets finished.

		__NOTE::
				deferred functions are stacking-up which is means that 'dd2' gonna be
				triggered first then 'dd'

		__NOTE::
				deferred functions based on function's scope.
	*/

	defer dd2()
}

// when it returning multiple values, the returned values's types should be in (type, type)
func doSomThing(a int, b int) (int, error) {
	return a * b, nil
}

func dd() {
	fmt.Println("this function gonna be called only when the hosted function gets closed/finished.")
}

func dd2() {
	fmt.Println("defers function `STACKING UP`")
}
