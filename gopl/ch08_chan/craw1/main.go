package main

import (
	"fmt"
	"go_demo/gopl/ch05_func/links"
	"log"
	"os"
)

func craw(url string) []string {
	fmt.Println(url)
	links, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return links
}

func main() {
	workList := make(chan []string)
	go func() {
		workList <- os.Args[1:]
	}()
	seen := make(map[string]bool)
	for list := range workList {
		fmt.Println("AAA")
		for _, url := range list {
			if !seen[url] {
				seen[url] = true
				go func(u string) {
					workList <- craw(u)
				}(url)
				fmt.Println("BBB")
			}
		}
	}
}

