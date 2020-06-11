package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("contex:\t", ctx) // contex:	 context.Background

	fmt.Println("context err:\t", ctx.Err()) // context err:	 <nil>

	fmt.Printf("context type: \t%T\n", ctx) // context type: 	*context.emptyCtx
}