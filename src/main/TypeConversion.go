package main

import "fmt"

type hotdog int

func typeDemo() {
	fmt.Println("Type demo")

	var x hotdog = 10
	fmt.Println(x)
	var y int = int(x)
	fmt.Println(y)
}
