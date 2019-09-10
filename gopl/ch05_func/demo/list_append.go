package main

import "fmt"

func main() {
	s := []string {"aaa", "bbb", "ccc", "ddd"}
	s1 := []string {"AAA", "BBB", "CCC", "DDD"}
	s = append(s, s1 ...)
	for _, t := range s {
		fmt.Println(t)
	}
}
