package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	lisener, err := net.Listen("tcp", "192.168.124.11:8888")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lisener.Accept()
		if err != nil {
			fmt.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	sc := bufio.NewScanner(c)
	cwd := "."
CLOSE:
	for sc.Scan() {
		args := strings.Fields(sc.Text())
		cmd := args[0]
		switch cmd {
		case "close":
			break CLOSE
		case "ls":
			if len(args) < 2 {
				ls(c, cwd)
			} else {
				path := args[1]
				if err := ls(c, path); err != nil {
					fmt.Fprintln(c, err)
				}
			}
		case "cd":
			if len(args) < 2 {
				fmt.Fprintln(c, "cd need a path")
			} else {
				cwd += "/" + args[1]
			}
		case "get":
			if len(args) < 2 {
				fmt.Fprintln(c, "get need a file name")
			} else {
				fielname := args[1]
				data, err := ioutil.ReadFile(fielname)
				if err != nil {
					fmt.Fprint(c, err)
				}
				fmt.Fprintf(c, "%s\n", data)
			}
		case "send":
			filename := args[1]
			f, err := os.Create(filename)
			if err != nil {
				fmt.Fprint(c, err)
			}
			defer f.Close()

			k, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Fprint(c, err)
			}

			var texts string
			for i := 0; i < k && sc.Scan(); i++ {
				texts += sc.Text() + "\n"
			}
			texts = strings.TrimSuffix(texts, "\n")
			fmt.Fprint(f, texts)
		}
	}
}

func ls(w io.Writer, path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Fprintf(w, "%s\n", file.Name())
	}
	return nil
}
