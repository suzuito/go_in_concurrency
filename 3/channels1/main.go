package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 100
	g := sync.WaitGroup{}
	g.Add(n)
	var ch1 chan int = make(chan int)
	var done chan int = make(chan int)
	go func() {
		// Writer
		for i := 0; i < n; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func(done chan<- int) {
		// Reader
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			fmt.Printf("I recieved %d\n", i)
		}
		close(done)
	}(done)
	_, ok := <-done
	fmt.Println(ok)
}
