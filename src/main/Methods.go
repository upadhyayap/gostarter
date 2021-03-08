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

// Attaches speak function to SecretAgent type i.e any value of type secret agent can call the function speak
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- secretAgent speaking")
}
