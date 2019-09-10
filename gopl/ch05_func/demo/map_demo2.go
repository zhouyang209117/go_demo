package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 1113
	for k, v := range m {
		fmt.Printf("k=%s,v=%d\n", k, v)
	}
	if value, ok := m["fdsakfa"]; ok {
		fmt.Println(m["a"])
		fmt.Println(value)
	}
}
