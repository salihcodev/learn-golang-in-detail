package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

type ByName struct{ Organs }
type ByWeight struct{ Organs }

func (s Organs) Len() int {
	return len(s)
}

func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}
func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {

	s := []Organ{{"Suzy", 845}, {"Rid", 345}, {"Fade", 120}, {"Mai", 1887}}

	sort.Sort(ByName{s})
	fmt.Println("ByName:", s)

	sort.Sort(ByWeight{s})
	fmt.Println("ByWeight:", s)

}
