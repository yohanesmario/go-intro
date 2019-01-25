package main

import (
	"fmt"
	"time"
)

func main() {
	worker := func() {
		fmt.Println("a")
	}
	go worker()
	go worker()
	go worker()

	time.Sleep(time.Second * 5)
}
