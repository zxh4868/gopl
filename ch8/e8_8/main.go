package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listien, err := net.Listen("tcp", "192.168.124.11:8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listien.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	text := make(chan string)
	go func() {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			text <- input.Text()
		}
	}()
	for {
		select {
		case <-time.After(10 * time.Second):
			return
		case x := <-text:
			go echo(conn, x, 1*time.Second)
		}
	}

}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))

}
