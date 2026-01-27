package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from goroutine")
}

func main_exerciseOne() {
	wg.Go(hello)
	wg.Wait()
}
