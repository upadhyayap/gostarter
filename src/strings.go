package main

import "fmt"

func main() {
	s := "Hello Go"
	s = "Hello"

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i])
	}

	for index, char := range s {
		fmt.Println(index, "-----", char)
	}

}
