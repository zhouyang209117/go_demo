package main

import "fmt"

func main() {
	m := make(map[string]bool)
	if m["aaaa"] {
		fmt.Println("AAA")
	} else {
		fmt.Println("BBB")
	}
}
