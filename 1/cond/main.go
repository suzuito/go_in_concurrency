package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	l := &sync.Mutex{}
	fmt.Println(l)
	l.Lock()
	c := sync.NewCond(l)
	f1 := func() {
		for range time.Tick(time.Second * time.Duration(rand.Intn(3))) {
			fmt.Println("f1")
			c.Signal()
		}
	}
	f2 := func() {
		for range time.Tick(time.Second * time.Duration(rand.Intn(3))) {
			fmt.Println("f2")
			c.Signal()
		}
	}
	// f3 := func() {
	// 	fmt.Println("f3 begins")
	// 	c.Wait()
	// 	fmt.Println("f3 ends")
	// }

	go f1()
	go f2()
	// go f3()

	for {
		fmt.Println("Waiting for signal ...")
		c.Wait()
		fmt.Println("Got")
	}
}
