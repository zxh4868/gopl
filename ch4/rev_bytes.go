package main

func reverse_byte(b []byte) {
	n := len(b)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func rev_bytes(s []byte, n int) {
	reverse_byte(s)
	reverse_byte(s[:n])
	reverse_byte(s[n-1:])
}

func main() {

}
