package main

import "fmt"

func getFullName() (string, string, string) {
	return "go", "lang", "dart"
}

func main() {
	fistName, midleName, _ := getFullName()
	fmt.Println(fistName)
	fmt.Println(midleName)
}
