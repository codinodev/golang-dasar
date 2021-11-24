package main

import "fmt"

type Backlist func(string) bool

func registerUser(name string, baklist Backlist) {
	if baklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	backlist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", backlist)
	registerUser("novan", backlist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("novan", func(name string) bool {
		return name == "root"
	})
}
