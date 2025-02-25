package main

import "fmt"

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)

func main() {
	var f float64 = 3 + 0i
	fmt.Printf("%T\n", f)
}
