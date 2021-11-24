package main

import "fmt"

func main() {

	name := "novan"
	if name == "novan" {
		fmt.Println("hello novan")
	} else if name == "joko" {
		fmt.Println("hi Joko")
	} else {
		fmt.Println("tidak ada")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama terlalu pendek")
	}

}
