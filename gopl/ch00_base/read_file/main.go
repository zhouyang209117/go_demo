package main

import (
	"bufio"
	"log"
	"os"
	"path"
)

func main() {
	filepath := "/Users/zhouyang3/go_project/go_demo/gopl/ch00_base/read_file/a.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	//pre := "/Users/zhouyang3/software/spark-2.4.3-bin-hadoop2.7/jars"
	pre := "/Users/zhouyang3/software/hbase-1.4.10/lib"
	result := ""
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += path.Join(pre, scanner.Text()) + ","
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	print(result)
}