package main

import "fmt"

func sayHello(firtsName string, lastName string) {
	fmt.Println("Hello", firtsName, lastName)
}

func main() {
	sayHello("kiki", "golang")
}
