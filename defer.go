package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println(result)

}

func main() {
	runApplication(10)
}
