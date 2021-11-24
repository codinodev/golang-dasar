package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "go",
		"addres": "google",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["addres"])
	person["title"] = "programer"
	fmt.Println(person)

	delete(person, "title")
	fmt.Println(person)

}
