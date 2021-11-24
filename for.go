package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke ", counter)

	}

	slice := []string{"novan", "golang", "dart"}
	//metode range
	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	person := make(map[string]string)
	person["name"] = "golang"
	person["title"] = "programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
