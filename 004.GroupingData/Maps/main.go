package main

import "fmt"

func main() {
	m := map[string]int{
		"Darko": 25,
		"Marko": 30,
	}

	fmt.Println(m)
	fmt.Println(m["Darko"])

	v, ok := m["Damir"]
	fmt.Println(v)  // value is 0
	fmt.Println(ok) // false

	if v, ok := m["Samir"]; ok {
		fmt.Println("Value exsists", v)
	} else {
		fmt.Println("Value do not exsists")
	}

	m["Marin"] = 32

	for key, value := range m {
		fmt.Println(key, value)
	}

	delete(m, "Darko")

	fmt.Println(m)

	someKey := "SomeKey"

	if v, ok := m[someKey]; ok {
		fmt.Println("Successfully deleted ", v)
		delete(m, someKey)
	}
}
