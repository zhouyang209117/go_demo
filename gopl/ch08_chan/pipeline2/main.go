package main

import (
	"fmt"
	"time"
)

func main() {
	natuals := make(chan int)
	squares := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			natuals <- i
		}
		close(natuals)
	}()

	go func() {
		for x := range natuals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
		time.Sleep(100 * time.Millisecond)
	}
}
