package main

import (
	"fmt"
	"time"
)

func ping(ch chan string) {
	ch <- "ping"
}

func main() {
	ch := make(chan string)
	go ping(ch)

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Buffer length:", len(ch))
}
