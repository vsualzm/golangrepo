package main

import "fmt"

func main() {

	var name string

	name = "ilham maulana"
	fmt.Println(name)

	name = "sugianto"
	fmt.Println(name)

	var age = 30
	fmt.Println(age)

	kota := "bandung"
	nickname := "vsualzm"
	umur := 21

	fmt.Println("--FORM--")
	fmt.Println(nickname)
	fmt.Println(kota)
	fmt.Println(umur)

	// menggabungkan variable dalam 1 var

	var (
		firstName = "ilham"
		lastName  = "maulana"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	var (
		namaBarang = "vans OLD SKOOL"
		ukuran     = 41
	)
	fmt.Println(namaBarang)
	fmt.Println(ukuran)
}
