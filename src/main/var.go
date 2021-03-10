package main

import "fmt"

//keyword identifier type
var foo int

const st = 10

const (
	a = 10
	b = 20
)

// iota: automatically resets value to 0 on next assignment
const (
	c = iota //0
	d = iota //1
	e = iota //2
)

func varDemo() {
	// var bar = 10
	// fmt.Println(foo, bar)

	// foo = 42

	// fmt.Println(foo)

	// fmt.Println(a)
	// fmt.Println(b)

	fmt.Println(c)
	fmt.Println(d)

	// resets the value whenever it is used
	const x = iota //x = 0

	fmt.Println(x)

}
