package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(conn net.Conn, content string, delay time.Duration) {
	fmt.Fprintf(conn, "\t%s\n", content)
	time.Sleep(delay)
	fmt.Fprintf(conn, "\t%s\n", strings.ToUpper(content))
	time.Sleep(delay)
	fmt.Fprintf(conn, "\t%s\n", strings.ToLower(content))
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		echo(conn, input.Text(), 1 * time.Second)
	}
}

func main() {
	listern, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listern.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}