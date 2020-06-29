package main

import (
	"Go-fundamentals/024.Testing/03.cat/mystr"
	"fmt"
	"strings"
)

const s = "Lorem Ipsum is simply dummy text of the printing and typesetting industry." +
	" Lorem Ipsum has been the industry's standard dummy text ever since the 1500s," +
	" when an unknown printer took a galley of type and scrambled it to make a type specimen book." +
	" It has survived not only five centuries, but also the leap into electronic typesetting, " +
	"remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing " +
	"Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}
