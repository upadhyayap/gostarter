package main

import "fmt"

// package level scopping, avoid package scoped variables
var x int

func closureDemo() {
	fmt.Println("Closure Demo")

	{
		y := 100
		fmt.Println(y)
	}
	//fmt.Println(y) // Error y will be undefined here because it's scope is lmited to above closure scope
	inc := incrementor()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

func incrementor() func() int {
	var val int
	return func() int {
		val++
		return val
	}
}
