package main

import (
	"fmt"
	"time"
	// "strconv"
	// "time"
)

func main() {
	input := make(chan int)
	// output := make(chan int)

	// feed numbers into input channel
	go func() {
		i := 1
		for {
			input <- i
			i++
			time.Sleep(time.Second * 1)
		}
	}()

	// print from input channel
	for {
		fmt.Println(<-input)
	}

	// get number from input channel, add 1,
	// pass it to output channel
	// go func() {
	// 	for {
	// 		p := <-input
	// 		output <- p + 1
	// 	}
	// }()

	// print from output channel
	// go func() {
	// 	for {
	// 		fmt.Println("Thread-1: " + strconv.Itoa(<-output))
	// 	}
	// }()

	//loop forever on main thread
	for {
	}
}
