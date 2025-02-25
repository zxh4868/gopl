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
	messages = make(chan string)
)

func main() {

	listener, err := net.Listen("tcp", "192.168.124.11:8080")
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

func breadcaster(){
	clients := make(map[client]bool)

	for {
		select {
		case msg := <- messages:
			for cli := range clients{
				cli.ch <- msg
			}
		case cli := <- entering:
			clients[cli] = true
		case cli := <- leaving:
			delete (clients, cli)
			close(cli.ch)
		}
	}
}

const timeout = 5 * time.Minute

func handleConn(conn net.Conn) {
	// 1. 获取名称
	var who string
	fmt.Fprintln(conn, "please input you name: ")
	fmt.Fscan(conn, &who)
	ch := make(chan string)
	go clientWriter(conn, ch)

	timer := time.NewTimer(timeout)
	go func() {
		for {
			<-timer.C
			conn.Close()
		}
	}()

	messages <- who + "has arrived"
	entering <- client{who, ch}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + " : " + input.Text()
		timer.Reset(timeout)
	}
	messages <- who + " has left"
	leaving <- client{who, ch}
	conn.Close()

}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
