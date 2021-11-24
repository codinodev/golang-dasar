package main

import "fmt"

func main() {
	var (
		name1 = "golang"
		name2 = "dart"

		result bool = name1 == name2

		value1 = 100
		value2 = 200
	)
	fmt.Println(result)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
