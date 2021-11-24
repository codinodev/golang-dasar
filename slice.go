package main

import "fmt"

func main() {

	month := [...]string{
		"januari",
		"febary",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	slice1 := month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// month[5] = "diubah"
	// fmt.Println(slice1)

	slice1[0] = "mei lagi"
	fmt.Println(month)

	slice2 := month[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "golang")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(month)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "go"
	newSlice[1] = "dart"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}
