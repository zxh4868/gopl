package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "192.168.124.11:8888")

	if err != nil {
		log.Fatal(err)
	}

	go broadercaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string) // 所有接收到的客户的消息
)

func broadercaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-message:
			// 把接收到的消息广播给所用客户端
			for clt := range clients {
				clt <- msg
			}
		case clt := <-entering:
			clients[clt] = true
		case clt := <-leaving:
			delete(clients, clt)
			close(clt)
		}
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	ch := make(chan string)
	go clientWriter(c, ch)

	who := c.RemoteAddr().String()
	ch <- "You are " + who
	message <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(c)
	for input.Scan() {
		message <- who + " : " + input.Text()
	}

	leaving <- ch
	message <- who + " has left"
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
