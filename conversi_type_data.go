package main

import "fmt"

func main() {
	//konversi tyoe data 1
	var (
		nilai32 int32 = 100000
		nilai64 int64 = int64(nilai32)
		nilai8  int8  = int8(nilai32)
	)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//konversi tyoe data 2
	var (
		name           = "go"
		e       uint8  = name[0]
		eString string = string(e)
	)
	fmt.Println(name)
	fmt.Println(eString)

}
