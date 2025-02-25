package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	name string
	ch   chan string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	massages = make(chan string)
)

func main() {

	listener, err := net.Listen("tcp", " 192.168.124.11")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

const timeout = 5 * time.Minute

func handleConn(conn net.Conn) {
	ch := make(chan string) // 向客户发送消息的通道

	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		conn.Close()
	}()

	go writerClient(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "You are" + who
	entering <- client{who, ch}
	massages <- who + "has arrived"

	input := bufio.NewScanner(conn)

	for input.Scan() {
		massages <- who + " : " + input.Text()
		timer.Reset(timeout)
	}
	leaving <- client{who, ch}
	conn.Close()
}

func writerClient(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
