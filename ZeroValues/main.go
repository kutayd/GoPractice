package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2
	fmt.Println(a, b)

	// a = b => compile type error
	// The type of a variable cannot be changed once declared

	a = int(b)
	fmt.Println(a, b)

	// GO Zero Values:
	// numeric: 0
	// bool: false
	// string: ""
	// pointer: nil

	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)

}
