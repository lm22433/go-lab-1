package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Printf("Hello from goroutine %d\n", i)
}

func main() {
	for i := 0; i < 5; i++ {
		go hello(i)
	}
	time.Sleep(1 * time.Second)
}
