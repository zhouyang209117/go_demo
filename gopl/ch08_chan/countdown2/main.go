package main

import (
	"fmt"
	"os"
	"time"
)

func lauch()  {
	fmt.Println("Lift off")
}

func main() {
	abort := make(chan bool)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- true
	}()
	select {
		case <-time.After(10 * time.Second):
		case <-abort:
			fmt.Println("Abort.")
			return
	}
	lauch()
}

