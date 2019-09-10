package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walkDir(rootDir string, fileSize chan<- int64)  {
	for _, entry := range dirents(rootDir) {
		if entry.IsDir() {
			walkDir(filepath.Join(rootDir, entry.Name()), fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func dirents(currentDir string) []os.FileInfo  {
	entries, err := ioutil.ReadDir(currentDir)
	if err != nil {
		fmt.Println("readDir error")
		return nil
	}
	return entries
}

func main() {
	flag.Parse()
	inputs := flag.Args()
	fileSize := make(chan int64)
	go func() {
		for _, rootDir := range inputs {
			walkDir(rootDir, fileSize)
		}
		close(fileSize)
	}()
	totalSize := int64(0)
	totalNum := int64(0)
	for size := range fileSize {
		totalSize += size
		totalNum += 1
	}
	fmt.Printf("total size=%d, total num=%d\n", totalSize, totalNum)
}