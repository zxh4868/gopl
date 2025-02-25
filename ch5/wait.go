package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 使用指数退避算法完成ping
func waitForServer(url string) error {
	const timeout = 60 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil // 连接成功

		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	if err := waitForServer("http://google.com"); err != nil {
		log.Fatalf("site is down: %v\n", err)
	}

}
