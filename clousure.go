package main

import "fmt"

func main() {
	name := "goalng"
	counter := 0
	increment := func() {
		name := "dart"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}
	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)

}
