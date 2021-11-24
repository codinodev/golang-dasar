package main

import "fmt"

func getNameFull() (firtsName string, midleName string, lastName string) {
	firtsName = "go"
	midleName = "lang"
	lastName = "dart"
	return
}

func main() {
	a, b, c := getNameFull()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
