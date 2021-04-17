package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
	channels block
	chan type is a type
	<- chan type is a type
	buffer channel
	Directional Chanels: what is the use of them
	Ranges
	select
	comma,Ok idiom
	FanIn
	FanOut
	context = when you have a process which starts a bunch of go routines and if you want to cancel them then context is used so that they do not leak
*/

func ChannelsDemo() {
	fmt.Println("Channels Demo")
}

func Deadlock() {
	c := make(chan int) // This will pass if we use a bufer channel

	c <- 42

	fmt.Println(<-c)
}

func SampleChannel() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func bufferChannel() {
	c := make(chan int, 1) // 1 value is allowed to sit in there

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func DirectionalChannel() {
	c1 := make(chan<- int)    // values can only be written into channel can not be retrived
	c2 := make(<-chan int, 2) // receive only channel, values can not be put into channel

	fmt.Printf("%T", c1)
	fmt.Printf("%T", c2)
}

// Using Channels

func usingChannels() {
	c := make(chan int)

	go fooChannel(c)

	barChannel(c) // This need not to be in a separate routine because <-c is a blocking call, it will block until foo does not puts a data in it

	fmt.Println("Exiting")

}

func fooChannel(c chan<- int) {
	fmt.Println("Adding on to the channel")

	c <- 42
}

func barChannel(c <-chan int) {
	fmt.Println("Received: ", <-c)
}

func RangeOverChannel() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	//pulls until channel is not closed
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Exiting")
}

func SelectDemo() {
	fmt.Println("Select From channel demo")

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("Exiting")
}

// converting general channel to receive only chanell
func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from even channel", v)
		case v := <-o:
			fmt.Println("from odd channel", v)
		case v := <-q:
			fmt.Println("from quit channel", v)
			return
		}
	}
	/* 	for {
		select {
		case v := <-e:
			fmt.Println("from even channel", v)
		case v := <-o:
			fmt.Println("from odd channel", v)
		case i, OK := <-q:
			// when channel has no value OK will be false
			if !OK {
				fmt.Println("quiting from chanel")
			} else {
				fmt.Println("from quit channel", i)
			}

			return
		}
	} */

	// Alternate Impl using command ok idiom

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
	close(q)
}

func add(a, b int) {
	fmt.Printf("%T", a)
}

func FanOutDemo() {
	c1 := make(chan int)
	c2 := make(chan int)
	go populate(c1)
	go FanOut(c1, c2)

	for v := range c2 {
		fmt.Println(v, "fetched from timeconsuming work")
	}

	fmt.Println("Exiting")
}

func populate(c chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("putting ", i, " on to channel")
		c <- i
	}

	close(c)
}

func FanOut(c1, c2 chan int) {
	var wg sync.WaitGroup
	// Adding throtteling
	const goroutines = 10
	wg.Add(goroutines) // only 10 routines will be launched because there are only 10
	for v := range c1 {
		//wg.Add(1)
		go func(v1 int) {
			c2 <- Timeconsumingwork(v1)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func Timeconsumingwork(data int) int {
	fmt.Println(data, "Time consuming work done")
	//fmt.println(data, "Time consuming work done")
	return data
}

func ContextDemo() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("received context done signal")
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 10)

	fmt.Println("Number of go routines", runtime.NumGoroutine())
	cancel()
	ctx.Done()
	fmt.Println("Send done signal and waiting for 2 sec")

	time.Sleep(time.Second * 2)

	fmt.Println("Number of go routines after cancel", runtime.NumGoroutine())
}
