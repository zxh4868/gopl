package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// 定义命令行参数
	port := flag.String("port", "8000", "port number")

	// 解析命令行参数
	flag.Parse()
	lisenter, err := net.Listen("tcp", "192.168.124.11:"+string(*port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lisenter.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
