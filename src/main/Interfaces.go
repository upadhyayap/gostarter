package main

import "fmt"

// This decleration tells us that any value which has function speak is also a type humen
// Every value implements a inbuilt type Empty interface interface{} so every value in go is also an empty interface type
type humen interface {
	speak()
}

func interfaceDemo() {
	// since secretAgent has speak function which means it is also a type humen along with secretAgent. A value in go can be of more than one type
	sa1 := secretAgent{
		person: person{
			first: "sardar",
			last:  "bhai",
			age:   34,
		},
		ltk: true,
	}

	p1 := person{
		first: "Rohit",
		last:  "Sharma",
		age:   34,
	}

	bar(sa1)
	bar(p1)
}

func bar(h humen) {
	fmt.Println(h)

	// special type of switch: way to implement factory pattern
	switch h.(type) {
	case person:
		fmt.Println("I am person", h.(person).first)
		break
	case secretAgent:
		fmt.Println("I am secret Agent", h.(secretAgent).first)
		break
	}
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- person speaking")
}
