package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
Concurrency vs paralllelism
paralllelism is an abilitty to run multiple tasks at the same time one of the
primary requirements of the paralllelism is multiple CPU's(multi core).
whereas concurrency is the design pattern to write a code in such a way that it can run in parallel

Wait groups

go routine
1. adding go before a function call makes it lunch it as a go routine
2. control flow does not need to wait for the routine to execute
3. when func main exits routines also exits

A go statment starts the execution of a function call as an independent thread or called go routine, within same address space.
Goroutines
*/

var wg sync.WaitGroup

func ConcurrencyDemo() {
	fmt.Println("Concurrency Demo")
	fmt.Println("Runtime props:")
	osprops()
	fmt.Println("End runtime props")

	// adding go will make the function run as a go routine
	// launcing a go routine
	wg.Add(1)
	go foocon()
	barcon()

	wg.Wait()
	fmt.Println("Number of go routines \t\t", runtime.NumGoroutine())
}

func osprops() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.GOARCH)
}

func foocon() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func barcon() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func raceConditionDemo() {
	fmt.Println("no of CPUS are : ", runtime.NumCPU())
	fmt.Println("no of go routines are : ", runtime.NumGoroutine())

	counter := 0

	gs := 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("no of CPUS are : ", runtime.NumCPU())
	fmt.Println("no of go routines are : ", runtime.NumGoroutine())
	fmt.Println("counter value is: ", counter)
}

func mutexDemo() {
	fmt.Println("no of CPUS are : ", runtime.NumCPU())
	fmt.Println("no of go routines are : ", runtime.NumGoroutine())

	counter := 0

	gs := 100

	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()
			v := counter
			//runtime.Gosched()
			v++
			counter = v
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("no of CPUS are : ", runtime.NumCPU())
	fmt.Println("no of go routines are : ", runtime.NumGoroutine())
	fmt.Println("counter value is: ", counter)
}

func atomicDemo() {
	fmt.Println("no of CPUS are : ", runtime.NumCPU())
	fmt.Println("no of go routines are : ", runtime.NumGoroutine())

	var counter int64

	gs := 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("no of CPUS are : ", runtime.NumCPU())
	fmt.Println("no of go routines are : ", runtime.NumGoroutine())
	fmt.Println("counter value is: ", atomic.LoadInt64(&counter))
}
