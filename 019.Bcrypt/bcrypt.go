package main

import (
	"fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "password"
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pass) //
	fmt.Println(string(bs)) // Hashed password

	err = bcrypt.CompareHashAndPassword(bs, []byte(pass))

	if err != nil {
		fmt.Println("You can't login")
		return
	}

	fmt.Println("You're logged in")
}
