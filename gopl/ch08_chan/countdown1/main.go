package main

import (
	"fmt"
	"time"
)

func lauch() {
	fmt.Println("Life off")
}
func main() {
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	lauch()
}
