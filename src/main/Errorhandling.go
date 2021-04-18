package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
	Go does not have exceptions
	Go code uses error values to indicate an abnormal state
	error is a type in go
	creating own custom type error means implementing the error interface
	Always check for errors, do not leave it with _ variable
	log.pannicln() - with panic we can recover. deffered functions run in this case
	log.fatalIn - with fatal recovery is not possible whole program shuts down and deffered functions not run
	recover is only useful inside deferred functions
*/
func ErrorHandlingDemo() {
	n, err := fmt.Println("Error handling demo")

	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	/* var name string

	fmt.Println("what is your name")

	_, nameErr := fmt.Scan(&name)
	if nameErr != nil {
		panic(nameErr) // can be recovered by using recover
	}

	fmt.Println(name)

	f, fileError := os.Create("name.txt")
	defer f.Close()

	if fileError != nil {
		panic(fileError)
	}

	r := strings.NewReader("Anand")

	io.Copy(f, r) */

	f, err := os.Open("name.txt")
	logFile, _ := os.Open("log.txt")
	defer f.Close()
	defer logFile.Close()
	log.SetOutput(logFile) // sends the logs to log file
	if err != nil {
		//fmt.Println("Error happend", err)
		log.Println("error happened", err)

		//log.Fatalln(err)
	}
}

func recoverDemo() {
	f()
	fmt.Println("returned nirmally from f")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.") // This will not be called because there was panic in g() but above defer func will be called
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("negative number")
		// return 0, fmt.Errorf("Error", f) // formatting error
	}

	return 42, nil
}

func CustomErrorDemo() {
	val, err := CustomErrorType(2)

	if err != nil {
		fmt.Println(err.Error())
	}

	/* if reflect.TypeOf(err) == reflect.TypeOf(userError) {

	} */ // how to compare error of which type

	fmt.Println(val)
}

func CustomErrorType(val int) (int, error) {
	if val == 1 {
		return val, nil
	}
	return 0, &userError{"Anand", 32}
}

// ceating custom error

type userError struct {
	name string
	age  int
}

func (err *userError) Error() string {
	return "Error is " + err.name + strconv.Itoa(err.age)
}
