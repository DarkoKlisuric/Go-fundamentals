package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// WRITING
	/*
		f, err := os.Create("names.txt")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()

		r := strings.NewReader("Darko Klisurić")

		io.Copy(f, r)
	*/
	// READING
	f, err := os.Open("names.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	bs, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
