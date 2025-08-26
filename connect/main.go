package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:9000")
	defer conn.Close()
	conn.Write([]byte("Hola"))
	//io.Copy
	//imprimirlo
	//mandarlo directamtne a la salida con estardar out
	io.Copy(os.Stdout, conn)
}
