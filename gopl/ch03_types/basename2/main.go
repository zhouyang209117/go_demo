package main

import (
	"fmt"
	"strings"
)

func basename(filePath string) string {
	i := strings.LastIndex(filePath, "/")
	filePath = filePath[i + 1:]
	i = strings.LastIndex(filePath, ".")
	filePath = filePath[:i]
	return filePath
}

func main() {
	s := "aaa/bbb/ccc/ddd中国.txt"
	fmt.Println(basename(s))
}
