package main

import "fmt"

func main() {

	//konversi tipedata
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	var nilai32a int32 = 2010
	var nilaiFloat float32 = float32(nilai32a) * 21.21

	fmt.Println(nilaiFloat)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "ilham"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)

	var kata = "kata"
	var kata3 = "cinta"

	fmt.Println(kata + kata3)

}
