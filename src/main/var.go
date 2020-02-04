package main

import "fmt"

var foo int

const st = 10

const (
	a = 10
	b = 20
)

const (
	c = iota
	d = iota
)

func varDemo() {
	var bar = 10
	fmt.Println(foo, bar)

	foo = 42

	fmt.Println(foo)

	fmt.Println(a)
	fmt.Println(b)

}
