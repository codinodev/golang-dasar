package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("error dengan message: ", message)
	}

	fmt.Println("Aplikasi selesai")

}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(false)
	fmt.Println("eko")
}
