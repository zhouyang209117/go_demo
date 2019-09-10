package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Print(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("mustCopy")
	mustCopy(os.Stdout, conn)
}