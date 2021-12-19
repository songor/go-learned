// go test -v -cover -short -bench .
// go test -run x -bench .
package main

import "testing"

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshall("post.json")
	}
}
