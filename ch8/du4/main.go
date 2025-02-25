package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// 用于限制目录并发数量的计数信号量
var sema = make(chan struct{}, 20)

func walkDir(dirname string, n *sync.WaitGroup, fileSzies chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dirname) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dirname, entry.Name())
			go walkDir(subdir, n, fileSzies)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du3: %v on walkdir\n", err)
				continue
			}
			fileSzies <- info.Size()
		}
	}
}

func dirents(dir string) []os.DirEntry {
	defer func() { <-sema }()

	select {
	case sema <- struct{}{}:
	case <-done:
		return nil // 取消
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %v on dients\n", err)
		return nil
	}
	return entries
}

func main() {
	flag.Parse()

	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var n sync.WaitGroup
	fileSizes := make(chan int64)
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	// Print the results periodically.
	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			for range fileSizes {
				// Do nothing.
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.2f GB\n", nfiles, float64(nbytes)/1e9)
}
