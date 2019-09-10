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
		log.Print(err)
		return nil
	}
	return links
}
func main() {
	worklist := make(chan []string)
	unseen := make(chan string)
	go func() {
		worklist <- os.Args[1:]
	}()
	for i := 0; i < 2; i++ {
		go func() {
			for l := range unseen {
				fmt.Println("AAA")
				links := craw(l)
				fmt.Printf("len=%d\n", len(links))
				fmt.Println("BBB")
				for _, l1 := range links {
					fmt.Printf("url=%s\n", l1)
				}
				//worklist <- links
				fmt.Println("CCC")

				go func() {
					worklist <- links
				}()
			}
		}()
	}
	seen := make(map[string]bool)
	for list := range worklist {
		fmt.Println("DDDD")
		for _, url := range list {
			if !seen[url] {
				seen[url] = true
				unseen <- url
			} else {
				log.Printf("%s processed\n", url)
			}
		}
	}
}

