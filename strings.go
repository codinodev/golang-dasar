package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("novan fatkhurrohman", "novan"))
	fmt.Println(strings.Split("novan fatkhurrohman", " "))
	fmt.Println(strings.ToLower("novan fatkhurrohman"))
	fmt.Println(strings.ToUpper("novan fatkhurrohman"))
	fmt.Println(strings.ToTitle("novan fatkhurrohman"))
	fmt.Println(strings.Trim(" novan fatkhurrohman ", " "))
	fmt.Println(strings.ReplaceAll("novan novan novan", "novan", "golang"))
}
