package main

import "fmt"

// function decleration ; Func (r Receiver) Identifier(parameters) (int,bool)) {....}
// function are first class citizen in go. you can assign a function type and return a function type
func funcDemo() {
	fmt.Println("functions Demo")

	xi := []int{1, 2, 3, 4}
	//if you notice unfurlSlice method has variadic parameter but we are passin a slice.
	//unfurlSlice(xi)
	//Here we are unfurling slice and passing it to unfurlSlice function.
	unfurlSlice(xi...)

	// Anonymous function, similar to IIFE in javascript

	func(x int, y bool) {
		fmt.Println("Anonymous function execution", x)
	}(10, false)

	// function expression. func() is an internal type in go
	f := func() {
		fmt.Println("This is function expression")
	}
	f()
	fmt.Printf("%T\n", f) // func()

	// Returning a function
	f1 := returnFunc()
	f1()

	// callbacks
	vals := []int{1, 2, 3, 4}
	sumFunc := sum
	sumEven := evenOp(sumFunc, vals...)
	//sum1 := evenOp(sum, vals...) // function ref can be directly passed
	fmt.Println(sumEven)

	x, _ := fooMulti("anand")

	fmt.Println(x)
}

func fooSingle(s string) string {
	return "anand"
}

// function returning two values: yes in go function can return two values
func fooMulti(s string) (string, bool) {
	return "anand", true
}

// variadic parameters

func fooVariadic(x ...int) {
	fmt.Printf("%T\n", x) // return will be Slice int []int
}

// 0 or more parameters of type int, variadic paramter
func unfurlSlice(x ...int) int {
	fmt.Printf("%T\n", x) // slice int
	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum
}

func returnFunc() func() {
	f := func() {
		fmt.Println("Hello from returning func")
	}

	return f
}

func sum(x ...int) int {
	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum
}

func evenOp(op func(x ...int) int, vi ...int) int {
	var evenVals []int

	for _, v := range vi {
		if v%2 == 0 {
			evenVals = append(evenVals, v)
		}
	}

	return op(evenVals...)
}
