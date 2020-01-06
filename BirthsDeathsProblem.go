package main

//Given list of births and deaths, determine years where most people were alive
//https://stackoverflow.com/questions/32896439/given-list-of-births-and-deaths-determine-years-where-most-people-were-alive
//https://www.presidentsusa.net/birth.html
import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Name  string
		Born  int
		Death int
	}{
		// sorted by birth year
		{"John", 1735, 1826},   //Oct 30, 1735	July 4, 1826
		{"George", 1732, 1799}, //Feb, 22, 1732	Dec 14, 1799 },
		{"Thomas", 1743, 1826}, //Apr 13, 1743	July 4, 1826
		{"James", 1751, 1836},  //Mar 16, 1751	June 28, 1836
		{"James", 1758, 1832},  //Apr 28, 1758 	Westmoreland Co., Va. 	July 4, 1831 	New York, New York
		{"John", 1767, 1848},   //July 11, 1767 	Quincy, Mass. 	Feb 23, 1848 	Washington, D.C.
		{"Andrew", 1767, 1845}, //Mar 15, 1767 	Waxhaws, No/So Carolina 	June 8, 1845 	Nashville, Tennessee
	}

	// Sort by Born, keeping original order or equal elements.
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Death < people[j].Death
	})
	fmt.Println(people)

	//i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	//if i < len(a) && a[i] == x {
	//	fmt.Printf("found %d at index %d in %v\n", x, i, a)
	//} else {
	//	fmt.Printf("%d not found in %v\n", x, a)
	//}
}
