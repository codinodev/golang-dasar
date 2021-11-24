package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "go"
	names[1] = "golang"
	names[2] = "dart"
	var values = [3]int{
		98, 80, 70,
	}
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values)
	fmt.Println(len(values))
	values[2] = 100
	fmt.Println(values)
}
