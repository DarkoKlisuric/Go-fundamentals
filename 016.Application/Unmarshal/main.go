package main

import (
	"encoding/json"
	"fmt"
)

type persons struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

func main() {
	s := `[{"FirstName":"James","LastName":"Bond","Age":32},
	      {"FirstName":"Miss","LastName":"Moneypenny","Age":27}]`

	bs := []byte(s)

	var people []persons

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All of the data", people)

	for i, v := range people {
		fmt.Println("\nPerson nubmer", i)
		fmt.Println(v.FirstName, v.LastName, v.Age)
	}
}
