package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// way to store data in an encrypted way
func bcryptDemo() {
	s := `password123`

	bs, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	err := bcrypt.CompareHashAndPassword(bs, []byte(s))

	fmt.Println(s)
	fmt.Println(string(bs))
	if err == nil {
		fmt.Println("correct password")
	}
}
