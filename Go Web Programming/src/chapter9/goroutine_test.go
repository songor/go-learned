package main

import (
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	print()
}

func TestGoPrint(t *testing.T) {
	goPrint()
	time.Sleep(1 * time.Millisecond)
}

// go test -run x -bench . -cpu 1
// go test -run x -bench . -cpu 2
func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print()
	}
}

func BenchmarkGoPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint()
	}
}
