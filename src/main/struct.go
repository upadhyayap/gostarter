package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//using one struct into another, this is also called embedded struct
type secretAgent struct {
	person
	ltk bool
}

func StructDemo() {
	fmt.Println("Struct Demo")

	p1 := person{
		first: "anand",
		last:  "upadhyay",
		age:   37,
	}

	sa1 := secretAgent{
		person: p1,
		ltk:    true,
	}

	sa2 := secretAgent{
		person: person{
			first: "sardar",
			last:  "khan",
			age:   34,
		},
		ltk: true,
	}

	fmt.Println(p1)
	// accessing fields using dot notation. Think of it as an object in C# or java
	fmt.Println(p1.first)

	fmt.Println(sa1.person.age)
	fmt.Println(sa1.ltk)

	fmt.Println(sa2.first)

	// Anonymous Structs
	an1 := struct {
		name   string
		weight int
	}{
		name:   "dog",
		weight: 25,
	}

	fmt.Println(an1.name)

}
