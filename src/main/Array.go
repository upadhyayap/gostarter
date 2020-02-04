package main

import (
	"fmt"
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
}
