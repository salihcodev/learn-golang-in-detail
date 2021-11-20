package main

import "fmt"

type Month struct {
	Week, Day int
}

type Year struct {
	Month
	Season string
}

/*
	some sort of specific String method created only for Month or Year.
	some sort of overriding string method but not actual overriding, coz it means that there is
	an inheritance. But it's not. (no overriding)
*/
func (m Month) String() string {
	return fmt.Sprintf("Week number: %d, Day number: %d !!", m.Week, m.Day)
}

// @1
//func (y Year) String() string {
//	return fmt.Sprintf("Week number: %d, Day number: %d, Current season: %s", y.Week, y.Day, y.Season)
//}

// PrintWeakNum @2
func PrintWeakNum(m Month) string {
	return fmt.Sprintf("%d", m.Week)
}

func main() {

	m := Month{18, 126}
	y := Year{Month{9, 63}, "Summer"}

	fmt.Println(m) // Week number: 18, Day number: 126
	fmt.Println(y) // Week number: 9, Day number: 63, Current season: Summer
	fmt.Println(PrintWeakNum(m)) // Worked; 18
	//fmt.Println(PrintWeakNum(y)) // Failed; Not the same type
	fmt.Println(PrintWeakNum(y.Month)) // Worked 9; But working if we're passing the Year's Month



	/*
		without @1 func, Year can use Month's function because we were embedded Month in Year.
		but there is a gotcha here, String() is a builtin method, and we are modifying its behavior,
		So if we created a custom function it only gonna be executing the passed type to this method see: PrintWeakNum @2
	*/

}
