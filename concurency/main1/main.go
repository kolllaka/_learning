package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	//go count("sheep")
	//go count("fish")

	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Wait()
	//time.Sleep(time.Second * 2)
	//fmt.Scanln()
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

}
