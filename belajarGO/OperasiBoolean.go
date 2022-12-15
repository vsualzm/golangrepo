package main

import "fmt"

func main() {
	// && : dan
	// || : atau
	// !  : kebalikan

	var nilaiA = 90
	var nilaiB = 80

	var lulusNilaiAkhirA bool = nilaiA > 80
	var lulusNilaiAkhirB bool = nilaiB > 70

	var lulus bool = lulusNilaiAkhirA && lulusNilaiAkhirB
	fmt.Println(lulus)

	var (
		biologi    = 90
		matematika = 80
		ips        = 90
	)

	totalPelajaran := biologi + matematika + ips
	rataRata := totalPelajaran / 3

	var kelulusan bool = rataRata == 86
	fmt.Println(rataRata)
	fmt.Println(kelulusan)

}
