package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// 回声服务器

func main() {
	listener, err := net.Listen("tcp", "192.168.124.11:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var number sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() {
		number.Add(1)
		go echo(c, input.Text(), 1*time.Second, &number)
	}
	go func() {
		number.Wait()
		tcp := c.(*net.TCPConn)
		tcp.CloseWrite()
	}()
}
