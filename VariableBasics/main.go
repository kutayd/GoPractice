package main

import "fmt"

func main() {
	var age int = 23
	fmt.Println("Age:", age)

	var name = "Kutay"                //go can initialize automatically
	fmt.Println("Your name is", name) // if i comment this line it will give me an compile error
	// because i didnt use the variable
	// to prevent it
	_ = name

	s := "Learning golang!" // this variable only works in block scope
	fmt.Println(s)
}
