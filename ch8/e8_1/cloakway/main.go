package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type server struct {
	name    string
	addr    string
	message string
}

func main() {
	servers := parseArg(os.Args[1:])

	for _, s := range servers {
		conn, err := net.Dial("tcp", s.addr)
		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()
		go func(ser *server) {
			for {
				sc := bufio.NewScanner(conn)
				for sc.Scan() {
					ser.message = sc.Text()
				}
				if err != nil {
					log.Fatal(err)
				}
				time.Sleep(1 * time.Second)
			}
		}(s)
	}
	for {
		for _, s := range servers {
			fmt.Printf("%s : %s\n", s.name, s.message)
		}
		fmt.Println("------------------------------------")
		time.Sleep(1 * time.Second)
	}
}

func parseArg(args []string) (servers []*server) {
	for _, arg := range args {
		s := strings.SplitN(arg, "=", 2)
		if len(s) != 2 {
			continue
		}
		servers = append(servers, &server{name: s[0], addr: s[1]})
	}
	return
}
