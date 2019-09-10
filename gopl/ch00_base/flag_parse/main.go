package main

import (
	"flag"
	"fmt"
)


func main() {
	flag.Parse()
	inputs := flag.Args()
	fmt.Println(inputs)
}