package main

import "fmt"

func stringDemo() {
	s := "Hello Go"
	s = "Hello"

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for index, char := range s {
		fmt.Println(index, "-----", char)
	}

	// there is no while go

	var x = 0

	for x > 10 {
		//do something
		x++
	}

}
