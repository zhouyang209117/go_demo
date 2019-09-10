package main

import "fmt"

func main() {
	m := make(map[string]bool)
	if m["aaa"] {
		fmt.Println("AAA")
	} else {
		fmt.Println("BBB")
	}
	m["aaa"] = true
	if m["aaa"] {
		fmt.Println("AAA")
	} else {
		fmt.Println("BBB")
	}
}
