package main

import "fmt"

func main() {
	// > Lebih dari
	// < Kurang dari
	// >= Lebih dari sama denngan
	// <= Kurang dari sama dengan
	// == sama dengna
	// != tidak sama dengan

	var name1 = "ilham"
	var name2 = "ilham"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}
