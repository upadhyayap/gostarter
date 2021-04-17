//Packages are a way to organize your code at one place, think of it as a container where all the functions live. It encourages code reusability. silier to
// namespaces in C# and package in java
//First line in any go program is the package
//No function can live outside of a package
//It’s path is relative to src directory

package main

import (
	"fmt"
)

// gets called for initialization before main func
func init() {
	fmt.Println("Init func called")
}

// Main function inside main package is the entry point for go program
func main() {
	fmt.Println("Starting")
	//ArrayDemo()
	// shorthandDemo()
	// stringDemo()
	// typeDemo()
	//varDemo()
	//mapsDemo()
	//StructDemo()
	//methodDemo()
	//funcDemo()
	//closureDemo()
	//stringDemo()
	//DeferDemo()
	//PointerDemo()
	//jsonDemo()
	//SortDemo()
	//bcryptDemo()
	//ConcurrencyDemo()
	//mutexDemo()
	//atomicDemo()
	//usingChannels()
	//RangeOverChannel()
	//SelectDemo()
	//FanOutDemo()
	ContextDemo()
}
