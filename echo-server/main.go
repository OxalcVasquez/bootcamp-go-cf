package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func main() {
	ln, _ := net.Listen("tcp", ":9000")
	defer ln.Close()
	for {
		c, _ := ln.Accept()
		handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	in := bufio.NewScanner(conn)
	log.Printf("New connection %s", conn.RemoteAddr())
	for in.Scan() {
		conn.Write(append(in.Bytes(), '\n'))
	}
	time.Sleep(5 * time.Second)
}
