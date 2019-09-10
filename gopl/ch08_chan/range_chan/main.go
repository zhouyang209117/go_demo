package main

import "fmt"

func main() {
	c := make(chan string, 4)
	c <- "aaa"
	c <- "bbb"
	c <- "ccc"
	c <- "ddd"
	close(c)
	for a := range c {
		fmt.Println(a)
	}
}