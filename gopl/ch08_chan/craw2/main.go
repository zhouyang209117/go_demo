package main

import (
	"fmt"
	"go_demo/gopl/ch05_func/links"
	"log"
	"os"
)

var tokens = make(chan bool, 20)

func craw(url string) []string {
	fmt.Println(url)
	tokens <- true
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)
	var n int
	n++
	go func() {
		workList <- os.Args[1:]
	}()
	seen := make(map[string]bool)
	for; n > 0; n-- {
		fmt.Printf("n=%d\n", n)
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(l string) {
					workList <- craw(l)
				}(link)
			}
		}
	}
}

