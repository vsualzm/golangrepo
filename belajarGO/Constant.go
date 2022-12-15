package main

import "fmt"

func main() {
	const data = "menjadi data"
	const data2 = 12

	fmt.Println("tidak bisa di ubah kembali")
	fmt.Println(data)
	fmt.Println(data2)

	const (
		kalimat1 = "sugardady"
		kalimat2 = "batuha"
	)

	fmt.Println(kalimat1)
	fmt.Println(kalimat2)
}
