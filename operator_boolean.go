package main

import "fmt"

func main() {
	var (
		nilaiAkhir = 80
		absensi    = 80

		lulusNilai   = nilaiAkhir > 80
		lulusAbsensi = absensi > 80

		lulus = lulusNilai && lulusAbsensi
	)
	fmt.Println(lulusNilai)
	fmt.Println(lulusAbsensi)
	fmt.Println(lulus)
	fmt.Println(nilaiAkhir >= 80 && absensi >= 80)
}
