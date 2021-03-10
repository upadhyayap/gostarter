package main

import "fmt"

func methodDemo() {
	fmt.Println("Methods Demo")

	sa1 := secretAgent{
		person: person{
			first: "sardar",
			last:  "bhai",
			age:   34,
		},
		ltk: true,
	}

	fmt.Println(sa1)

	sa1.speak()

}
