package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

var strSlice = make([]string, LIMIT)

const LIMIT = 1010

func init() {
	for i := 0; i < LIMIT; i++ {
		strSlice[i] = loremIpsum
	}
}

func BenchmarkConcationOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q string
		for _, v := range strSlice {
			q = q + v
		}
	}
	b.ReportAllocs()
}

func BenchmarkFmtSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q string
		for _, v := range strSlice {
			q = fmt.Sprintf(q, v)
		}
	}
	b.ReportAllocs()
}

func BenchmarkBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q bytes.Buffer
		// 预分配内存
		q.Grow(len(loremIpsum) * len(strSlice))
		for _, v := range strSlice {
			q.WriteString(v)
		}
		_ = q.String()
	}
	b.ReportAllocs()
}

func BenchmarkStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q strings.Builder
		// 预分配内存
		q.Grow(len(loremIpsum) * len(strSlice))
		for _, v := range strSlice {
			q.WriteString(v)
		}
		_ = q.String()
	}
	b.ReportAllocs()
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q []byte

		for _, v := range strSlice {
			q = append(q, v...)
		}
		_ = string(q)
	}
	b.ReportAllocs()
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q string
		q = strings.Join(strSlice, "")
		_ = q
	}
	b.ReportAllocs()
}
