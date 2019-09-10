package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	go func() {
		mustCopy(os.Stdout, conn)
		log.Print("Done")
		done <- true
	}()
	mustCopy(conn, os.Stdin)
	<- done
}

func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("mustCopy error:%v", err)
	}
}