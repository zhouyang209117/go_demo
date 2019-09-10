package main

import "fmt"

func func1() {
	a := make(chan int)
	go func() {
		a <- 1
	}()
	fmt.Println(<-a)
}


func main() {
	func1()
}