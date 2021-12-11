package main

import (
	"fmt"
)

/* ASSIGNMENT INFO
inputs:
	- acceleration
	- initial velocity
	- initial displacement
	- time

output:
	- calculated displacement

notes:
	* the program should compute after entering the time.
*/

func HandleEndUser(a, v1, s1, t *float64) {
	var _a, _v1, _s1, _t float64
	fmt.Println("** Enter inputs in floats or integers **")
	fmt.Print("Acceleration: ")
	_, err := fmt.Scanf("%g", &_a)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Initial Velocity: ")
	_, _err := fmt.Scanf("%g", &_v1)
	if err != nil {
		fmt.Println(_err)
	}

	fmt.Print("Initial Displacement: ")
	_, __err := fmt.Scanf("%g", &_s1)
	if err != nil {
		fmt.Println(__err)
	}

	fmt.Print("Time In Seconds: ")
	_, ___err := fmt.Scanf("%g", &_t)
	if err != nil {
		fmt.Println(___err)
	}

	*a = _a
	*v1 = _v1
	*s1 = _s1
	*t = _t
}

func DisplaceLow(a, v1, s1, t float64) float64 {
	// s = ½ a t2 + vot + so
	return (0.5*a)*(t*2) + (v1 * t) + s1
}

func GenDisplaceFn(a, v1, s1 float64) func(time float64) float64 {
	fn := func(t float64) float64 {
		s := DisplaceLow(a, v1, s1, t)

		return s
	}

	return fn
}

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64

	// get the inputs from the end-user:
	HandleEndUser(&acceleration, &initialVelocity, &initialDisplacement, &time)

	// the closure point:
	DisplaceAt := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	/*
		coz the initial value of numbers vars equal 0,
		then if it changed that means that the user is just entered the time

		NOTE: can just call `fmt.Println(DisplaceAt(time))`
		without checking if the `time` is zero or not

		coz it's initially going be called in the stack after calculate the displacement
	*/
	if time != 0 {
		fmt.Println("Displacement: ", DisplaceAt(time))
	}

	// examples
	//fmt.Println("Displacement: ", DisplaceAt(5)) // Displacement: 26
	//fmt.Println("Displacement: ", DisplaceAt(3)) // Displacement: 18
}
