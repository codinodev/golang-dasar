package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regexp = regexp.MustCompile(`e([a-z])o`)
	fmt.Println(regexp.MatchString("eko"))
	fmt.Println(regexp.MatchString("edo"))
	fmt.Println(regexp.MatchString("eKo"))

	fmt.Println(regexp.FindAllString("eko edo egi ego e1o eto", 10))
	seacrh := regexp.FindAllString("eko edo egi ego e1o eto", 1)
	fmt.Println(seacrh)
}
