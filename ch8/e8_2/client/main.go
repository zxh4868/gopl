package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.124.11:8888")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 在后台把连接中的内容发送到标准输出
	go mustCopy(os.Stdout, conn)

	sc := bufio.NewScanner(os.Stdin)
CLOSE:
	for sc.Scan() {
		args := strings.Fields(sc.Text())
		cmd := args[0]

		switch cmd {
		case "close":
			fmt.Fprint(conn, sc.Text())
			break CLOSE
		case "ls", "cd", "get":
			fmt.Fprintln(conn, sc.Text())
		case "send":
			if len(args) < 2 {
				log.Println("send need a file name")
			} else {
				filename := args[1]
				data, err := ioutil.ReadFile(filename)
				if err != nil {
					log.Println("read file err:", err)
				}
				fmt.Fprintf(conn, "%s %d\n", sc.Text(), countlines(data))
				fmt.Fprintf(conn, "%s", data)
			}
		}
	}

}

func countlines(data []byte) int {
	c := 0
	sc := bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan() {
		c++
	}
	return c
}

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}
