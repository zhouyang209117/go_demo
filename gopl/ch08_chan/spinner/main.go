package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, c := range`-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(delay)
		}
	}
}

func fibN(n int) int {
	if n < 2 {
		return n
	} else {
		return fibN(n - 1) + fibN(n - 2)
	}
}

func main()  {
	go spinner(100 * time.Millisecond)
	n := 45
	fmt.Printf("\nfibN(%d)=%d\n", n, fibN(45))
}
