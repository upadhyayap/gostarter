package main

import (
	"fmt"
)

func DeferDemo() {
	defer fmt.Println("world")
	defer fmt.Println("navin")

	fmt.Println("hello")
	aDefer()

}

func aDefer() {
	i := 0
	defer fmt.Println(i) // output will be 0 here, A deffered functions arguments is evaluated when the deffered function is evalualetd
	i++
	return
}

// defered functions are evaluated in LIFO order

func DeferLIFO() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // 4,3,2,1,0
	}
}

// deffered function may read and assign value to surrounding functions named return values

func DeferReturn() (i int) {
	defer func() {
		i++
	}()

	return 1
} // returns 1
