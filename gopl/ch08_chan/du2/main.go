package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

var verbose = flag.Bool("v", false, "aaa")

func onWalk(rootDir string, fileSize chan int64)  {
	for _, entry := range dirents(rootDir) {
		if entry.IsDir() {
			onWalk(path.Join(rootDir, entry.Name()), fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func dirents(rootDir string) [] os.FileInfo {
	info, err := ioutil.ReadDir(rootDir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return info
}

func main() {
	flag.Parse()
	inputs := flag.Args()
	if len(inputs) == 0 {
		inputs = []string{"."}
	}
	fileSize := make(chan int64)
	go func() {
		for _, currentDir := range inputs {
			onWalk(currentDir, fileSize)
		}
	}()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	totalSize := int64(0)
	totalNum := int64(0)
loop:
	for {
		select {
			case size, ok := <- fileSize:
				if !ok {
					break loop
				}
				totalSize += size
				totalNum += 1
			case <- tick:
				fmt.Printf("num=%d,size=%.1fGB\n", totalNum, float32(totalSize) / 1e9)
		}
	}
}