package main

import "fmt"

func shorthandDemo() {
	bytesWritten, _ := fmt.Println("Hello Go", 42, true) // _ to ignore the second paramete

	fmt.Println(bytesWritten)
}
