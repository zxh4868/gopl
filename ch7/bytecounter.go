package main

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// func main() {
// 	var c ByteCounter
// 	c.Write([]byte("hello"))
// 	fmt.Println(c)

// 	c = 0 // 重置计数器
// 	var s = []byte("hello")
// 	fmt.Fprintf(&c, "%s,nihao", s)
// 	fmt.Println(c)
// }
