package main

import (
	"fmt"
	"sort"
)

func ArrayDemo() {
	fmt.Println("Array Demo")

	var x [5]int

	fmt.Println(x)

	x[0] = 1

	fmt.Println(x, len(x))

	y := [5]int{1, 2, 3, 4, 5} // This initialization is called a composite literal
	fmt.Println(y)

	for index, val := range y {
		fmt.Println("value at index ", index, " is ", val)
	}

	fmt.Println(cap(y))

	fmt.Println("slicing an array")

	fmt.Print(y[1:])    // slice array from index 1 to end
	fmt.Println(y[1:3]) // will slice array from index 1 to 3 Not including end index

	// Slice has length and capacity

	SortDemo()
}

func SortDemo() {
	fmt.Println("sort Demo")

	s := []int{5, 4, 7, 2, 8, 19, 1}
	sort.Ints(s)

	fmt.Println(s)

	sort.Strings([]string{"a", "b"})

}

type SortByAge []Person

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func CustomSort() {
	p1 := Person{
		First: "Anand",
		Last:  "upadhyay",
		Age:   30,
	}

	p2 := Person{
		First: "Aditya",
		Last:  "Gaurav",
		Age:   32,
	}

	people := []Person{p1, p2}

	sort.Sort(SortByAge(people))

	fmt.Println(people)
}

func testingArray() []int {
	return []int{1, 2}
}

func avgArray(data []int) float64 {
	var sum float64

	for _, val := range data { // _ is index
		sum += float64(val)
	}

	return sum / float64(len(data))
}
