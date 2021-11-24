package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(addres *Address) {
	addres.Country = "indonesia"
}

func main() {
	//operator* *address
	//operator& &address
	//function new
	address1 := Address{"tegal", "Jawa Tengah", "indonesia"}
	//address4 := Address{"tegal", "Jawa Tengah", "indonesia"}
	address2 := &address1
	address3 := &address1
	address2.City = "slawi"

	*address2 = Address{"jakarta", "dki jakarta", "indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address4.City = "malang"
	fmt.Println(address4)

	alamat := Address{
		City:     "pemalang",
		Province: "jawa tengah",
		Country:  "",
	}
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}
