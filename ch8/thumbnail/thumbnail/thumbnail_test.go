// This file is just a place to put example code from the book.
// It does not actually run any code in gopl/ch8/thumbnail.

package thumbnail

import (
	"log"
	"os"
	"sync"
)

// makeThumbnails 生成指定文件的缩略图
func makeThumbnails(finenames []string) {
	for _, f := range finenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// NOTE: incorrect! 因为 goroutine 在 主goroutine 函数返回之前还没有执行完毕
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go ImageFile(f) // NOTE: ignoring errors
	}
}

// makeThumbnails3 并行生成指定文件的缩略图
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f)
	}

	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}

// makeThumbnails4 makes thumbnails for specified files in parallel.
// 它可以返回任意文件的错误。
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: incorrect: goroutine leak!当第一个非nil的错误被接受后，剩下的goroutine将被阻塞在发送到errors通道上
		}
	}

	return nil
}

// makeThumbnails5 makes thumbnails for the specified files in parallel.
// 她以任意顺序返回生成文件名
// 如果任何步骤出错就返回一个错误
// 函数返回不会强制终止子goroutine（除非是main），子 goroutine 会继续运行，直到完成自己的任务
// 通道不会被释放：只有当没有任何 goroutine 引用通道时，通道才会被垃圾回收
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

// makeThumbnails6 makes thumbfiles for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
