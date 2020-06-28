package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("err happend", err) // err happend open no-text-file.txt: no such file or directory
		log.Fatalln(err) // 2020/06/28 22:10:35 open no-text-file.txt: no such file or directory
		// panic(err)
	}

	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}
