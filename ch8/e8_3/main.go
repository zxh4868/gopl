package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.124.11:8888")
	tcp := conn.(*net.TCPConn)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan int)
	go func() {
		io.Copy(os.Stdout, conn)
		log.Panicln("done")
		done <- 1
	}()

	mustCopy(conn, os.Stdin)
	tcp.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
