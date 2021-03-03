package main

import "fmt"

func funcDemo() {
	fmt.Println("functions Demo")
}

func fooSingle(s string) string {
	return "anand"
}

// function returning two values: yes in go function can return two values
func fooMulti(s string) (string, bool) {
	return "anand", true
}
