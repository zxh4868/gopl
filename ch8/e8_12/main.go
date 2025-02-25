package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	name string
	msg  chan<- string
}

// 全局变量，用于记录服务器的状态
var entering = make(chan client)
var leaving = make(chan client)
var message = make(chan string)

func main() {
	listener, err := net.Listen("tcp", " 192.168.124.11")

	if err != nil {
		log.Fatal("server error!")
	}
	go broadercaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handdleConn(conn)
	}
}

func broadercaster() {

	clients := make(map[client]bool) // 所有连接的客户端

	for {
		select {

		case msg := <-message:
			for cli := range clients {
				cli.msg <- msg
			}

		case cli := <-entering:
			// 收集当前服务器上所有客户的名字
			var names []string
			for x := range clients {
				names = append(names, x.name)
			}
			cli.msg <- fmt.Sprintf("%d arrival %v\n", len(names), names)
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msg)
		}
	}

}

func handdleConn(conn net.Conn) {
	ch := make(chan string) // 对外发送客户消息的通道
	go writeClient(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are" + who
	message <- who + "has arrived"
	c := client{who, ch}
	entering <- c

	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + " : " + input.Text()
	}
	leaving <- c
	message <- who + "has left"
	conn.Close()
}

func writeClient(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
