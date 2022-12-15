package main

import "fmt"

func main() {
	// membuat alias tipe data yang sudah ada

	type hape string
	type status bool

	type jeniskelamin string
	type noNIK int

	var nohp hape = "2313123"
	var kondisihp status = true

	var no noNIK = 121212

	fmt.Println(no)
	fmt.Println(nohp)
	fmt.Println(kondisihp)
}
