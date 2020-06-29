package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Darko")

	if s != "Welcome user Darko" {
		t.Error("got", s, "expected", "Welcome user Darko")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Darko"))
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Darko")
	}
}

// go test -bench .
// . is all Benchmark
// in this case with one benchmark ->
// go test -bench Greet - func BenchmarkGreet..

/*
goos: linux
goarch: amd64
pkg: Go-fundamentals/024.Testing/02.benchmark/01/saying
BenchmarkGreet-12    	20000000	       113 ns/op
PASS
ok  	Go-fundamentals/024.Testing/02.benchmark/01/saying	2.380s
*/

// Running on 12 cores, 20 millions times in 2.3s