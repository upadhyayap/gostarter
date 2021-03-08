package main

import "fmt"

func DeferDemo() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
