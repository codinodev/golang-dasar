package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var novan Customer
	novan.Name = "novan golang"
	novan.Address = "Indonesia"
	novan.Age = 20

	fmt.Println(novan.Name)
	fmt.Println(novan.Address)
	fmt.Println(novan.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Cirebon",
		Age:     20,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "jakarta", 20}
	fmt.Println(budi)
}
