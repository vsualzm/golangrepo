package main

import "fmt"

func main() {

	const (
		semangka = 1000
		rambutan = 2000
		mangga   = 3000
	)

	var (
		dukuh  = 2000
		sirsak = 1500
		jambu  = 3500
	)

	totalBuah := semangka + rambutan + mangga
	totalBuah2 := dukuh + sirsak + jambu
	fmt.Println("hasil seluruh buah adalah:", totalBuah)
	fmt.Println("hasil seluruh buah adalah:", totalBuah2)

	var a = 10

	a += 10
	fmt.Println(a)

	a++
	fmt.Println(a)

	/// unary operator
	// ++ di tambah 1
	// -- di kurang 1
	// - negatif
	// + positive
	// ! metode kebalikan

	var z = -100
	var x = +100
	fmt.Println(z)
	fmt.Println(x)

}
