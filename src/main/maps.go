package main

import (
	"fmt"
)

func mapsDemo() {
	fmt.Println("Maps Demo")

	// declaring map which has key of type string and value of type int
	m := map[string]int{
		"anand":  32,
		"aditya": 41,
	}

	//map["anand"]type

	fmt.Println(m)
	fmt.Println(m["anand"])

	v, ok := m["sardar"]

	fmt.Println(v)
	fmt.Println(ok)

	m["anand"] = 49

	// checking if value exists in a map

	if val, ok := m["anand"]; ok {
		fmt.Println(val)
	}

	// looping over map using range
	for k, v := range m {
		fmt.Print(k, v)
	}

	delete(m, "aditya")
}
