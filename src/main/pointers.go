package main

import "fmt"

// pointing to some location in memory where value is stored.This makes everything in go is a pass by value
// & to get the pointre to the variable and * to get the value from the pointer type
// when to use pointers: when you have big chunk of data so instead of passing the data you can pass the address where data lives.
//
func PointerDemo() {
	fmt.Println("pointer Demo")

	// Basics
	a := 42

	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // & gives you the address

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // derefrencing an address. gives you the value stored at the address

	// getting address and dereferencing at the same time
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)
	// END Basics

	x := 0
	foop(x)
	fmt.Println(x)

	barp(&x)
	fmt.Println(x)
}

func foop(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func barp(z *int) {
	fmt.Println(*z)
	*z = 43
	fmt.Println(*z)
}
