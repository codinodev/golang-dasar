package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Merried() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method", man.Name)
}

func main() {
	novan := Man{"novan"}
	novan.Merried()
	fmt.Println(novan.Name)
}
