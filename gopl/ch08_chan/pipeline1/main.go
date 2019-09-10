package main

import (
	"fmt"
	"time"
)

func main() {
	natuals := make(chan int)
	squares := make(chan int)
	go func() {
		for i := 1; ; i++ {
			natuals <- i
		}
	}()
	go func() {
		for  {
			x := <- natuals
			squares <- x * x
		}
	}()
	for {
		fmt.Println(<- squares)
		time.Sleep(1 * time.Second)
	}
}
