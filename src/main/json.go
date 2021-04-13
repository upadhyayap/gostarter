package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func jsonDemo() {
	fmt.Println("Json Demo")
	io.WriteString(os.Stdout, "hello")
	marshell()
}

// Json marshell and unmarshell
func marshell() {
	p1 := Person{
		First: "Anand",
		Last:  "bhai",
		Age:   30,
	}

	people := []Person{p1}

	val, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(val))
	}

	//os.Stdout.Write(json.Marshal(p1))
}
