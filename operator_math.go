package main

import "fmt"

func main() {
	var (
		a = 10
		b = 10
		c = a + b
	)

	i := 10
	i += 10
	fmt.Println(c)
	fmt.Println(i)
	i++
	fmt.Println(i)
}
