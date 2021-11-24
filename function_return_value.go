package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "hello dart"
	} else {
		return "hello " + name
	}
}

func main() {
	result := getHello("eko")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
