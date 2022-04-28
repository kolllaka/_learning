package main

import (
	"fmt"
	"time"
)

func main() {
	// c := make(chan string, 2)
	// c <- "hello"
	// c <- "world!"

	// msg := <-c
	// fmt.Println(msg)

	// msg = <-c
	// fmt.Println(msg)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2000ms"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
		
		// fmt.Println(<-c1)
		// fmt.Println(<-c2)
	}

}
