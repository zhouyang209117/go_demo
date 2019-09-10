package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walkDir(rootDir string)  {
	for _, entry := range dirents(rootDir) {
		if entry.IsDir() {
			walkDir(filepath.Join(rootDir, entry.Name()))
		} else {
			fmt.Println(filepath.Join(rootDir, entry.Name()))
		}
	}
}

func dirents(fileDir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(fileDir)
	if err != nil {
		fmt.Println("read dir error")
		return nil
	}
	return entries
}

func main() {
	rootDir := "/Users/zhouyang3/tmp/samples"
	walkDir(rootDir)
}