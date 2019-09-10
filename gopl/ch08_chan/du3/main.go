package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"
)
var vFlags = flag.Bool("v", false, "aaa")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSize := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go onWalk(root, &n, fileSize)
	}
	go func() {
		n.Wait()
		close(fileSize)
	}()
	var tick <-chan time.Time
	if *vFlags {
		tick = time.Tick(500 * time.Millisecond)
	}
	var fileNum, totalSize int64
loop:
	for {
		select {
		case size, ok := <- fileSize:
			if !ok {
				break loop
			}
			fileNum += 1
			totalSize += size
		case <- tick:
			fmt.Printf("num=%d,size=%.1f\n", fileNum, float32(totalSize) / 1e9)
		}
	}
}

func onWalk(currentDir string, n *sync.WaitGroup, fileSize chan int64)  {
	defer n.Done()
	for _, entry := range direntry(currentDir) {
		if entry.IsDir() {
			n.Add(1)
			subDir := path.Join(currentDir, entry.Name())
			go onWalk(subDir, n, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)
func direntry(currentDir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() {
		<-sema
	}()
	entry, err := ioutil.ReadDir(currentDir)
	if err != nil {
		return nil
	}
	return entry
}
