package main

import "fmt"

func DeferDemo() {
	defer fmt.Println("world")
	defer fmt.Println("navin")

	fmt.Println("hello")
}
