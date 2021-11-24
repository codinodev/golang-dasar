package main

import "fmt"

func main() {
	name := "go"

	switch name {
	case "golang":
		fmt.Println("hello golang")
	case "go":
		fmt.Println("hello go")

	}
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama benar")
	}
}
