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

func walkRoot(root string, m *sync.WaitGroup) {
	defer func() {
		m.Done()
	}()
	var n sync.WaitGroup
	fileSizes := make(chan int64)
	n.Add(1)
	go walkDir(root, &n, fileSizes)

	go func() {
		n.Wait()
		close(fileSizes)
	}()
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles += 1
		nbytes += size
	}
	printDiskUsage(root, nfiles, nbytes)
}

func main() {
	flag.Parse()

	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}
	var m sync.WaitGroup
	for _, root := range roots {
		m.Add(1)
		go walkRoot(root, &m)
	}

	m.Wait()

	fmt.Println("=========================Done===========================")
}

func printDiskUsage(root string, nfiles, nbytes int64) {
	fmt.Printf("%s : %d files %.2f GB\n", root, nfiles, float64(nbytes)/1e9)
}
