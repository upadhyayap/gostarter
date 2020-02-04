package main

import "fmt"

func main() {
	var foo int = 10
	var b bool = true
	fmt.Printf("%T\n", foo)
	fmt.Println(b)

	// Custom type
	type hotdog int
	var bar hotdog = 43

	fmt.Println(bar)
	fmt.Printf("%T\n", bar)

	// Type conversion

	// foo = bar Not allowed because they both are of different type, we need to do type conversion to acheive this

	// bar is converted to int and now can be assigned to foo
	foo = int(bar)
	fmt.Println(foo)

}
