package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// 用于限制目录并发数量的计数信号量
var sema = make(chan struct{}, 20)

func walkDir(dirname string, n *sync.WaitGroup, fileSzies chan<- int64) {
	defer n.Done()
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
	defer func() {
		<-sema
	}()
	sema <- struct{}{}
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
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles += 1
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.2f GB\n", nfiles, float64(nbytes)/1e9)
}
