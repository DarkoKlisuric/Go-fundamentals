package mystr

import (
	"strings"
	"testing"
)

const s = "Lorem Ipsum is simply dummy text of the printing and typesetting industry." +
	" Lorem Ipsum has been the industry's standard dummy text ever since the 1500s," +
	" when an unknown printer took a galley of type and scrambled it to make a type specimen book." +
	" It has survived not only five centuries, but also the leap into electronic typesetting, " +
	"remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing " +
	"Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."

func BenchmarkCat(b *testing.B) {
	xs := strings.Split(s, " ")

	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs := strings.Split(s, " ")

	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}

// go test -bench .

/*
goos: linux
goarch: amd64
pkg: Go-fundamentals/024.Testing/03.cat/mystr
BenchmarkCat-12     	  100000	     35673 ns/op
BenchmarkJoin-12    	 1000000	      1841 ns/op
PASS
ok  	Go-fundamentals/024.Testing/03.cat/mystr	5.576s
*/