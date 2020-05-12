package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// THIS IS THE SAME
	fmt.Println("Hello world")
	fmt.Fprintln(os.Stdout, "Hello world")
	io.WriteString(os.Stdout, "Hello world\n")
}
