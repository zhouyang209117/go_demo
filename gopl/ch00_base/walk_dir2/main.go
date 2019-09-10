package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var cnt int
func processFile(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	fmt.Printf("%d.%s\n", cnt, path)
	cnt += 1
	return nil
}

func main() {
	rootDir := "/Users/zhouyang3/tmp/samples"
	cnt = 1
	filepath.Walk(rootDir, processFile)
}