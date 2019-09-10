package main

import (
	"fmt"
	"os"
)

func main() {
	fileDir := "/Users/zhouyang3/CLionProjects/Yun2Txt"
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		fmt.Println("not exist.")
	}

	filePath := "/Users/zhouyang3/PycharmProjects/compare/demo1.py"
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("%s exists.", filePath)
	} else {
		fmt.Printf("%s not exists.", filePath)
	}
}
