package main

import "fmt"

func main() {
	// tipe data array
	// di limit kalau di golang

	var names [4]string
	names[0] = "iyam"
	names[1] = "ganteng"
	names[2] = "sedunia"
	names[3] = "memek"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	var values = [3]int{
		90,
		90,
		80,
	}

	fmt.Println(values)
	fmt.Println(values[1])
	fmt.Println(len(values))

}
