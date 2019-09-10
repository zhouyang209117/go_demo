package main

import (
	"fmt"
	"sync"
	"time"
)

func longProcess(id int, wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("id=%d is done\n", id)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go longProcess(i, &wg)
	}
	wg.Wait()
	fmt.Println("total finished.")
}
