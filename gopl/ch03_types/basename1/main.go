package main

import "fmt"

func basename(filePath string) string {
	for i := len(filePath) - 1; i >= 0; i-- {
		if filePath[i] == '/' {
			filePath = filePath[i + 1:]
			break
		}
	}
	for i := len(filePath) - 1; i >= 0; i-- {
		if filePath[i] == '.' {
			filePath = filePath[:i]
			break
		}
	}
	return filePath
}

func main() {
	s := "aaa/bbb/ccc/ddd中国.txt"
	fmt.Println(basename(s))
}
