package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(conn net.Conn, content string, delay time.Duration)  {
	fmt.Fprintf(conn, "\t%s\n", content)
	time.Sleep(delay)
	fmt.Fprintf(conn, "\t%s\n", strings.ToUpper(content))
	time.Sleep(delay)
	fmt.Fprintf(conn, "\t%s\n", strings.ToLower(content))
	time.Sleep(delay)
}

func handleConn(conn net.Conn)  {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), 10 * time.Second)
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Panicf("accept error:%v", err)
		}
		go handleConn(conn)
	}
}