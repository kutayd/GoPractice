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

	car, cost := "Audi", 500000
	fmt.Println(car, cost)

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary float64
		firstN string
		lastN  string
	)
	fmt.Println(salary, firstN, lastN)

	var a, b, c int
	fmt.Println(a, b, c)

	// We use short declaration (:=) when we know the initial value
	// We use normal declaration (var) otherwise

	var i, j int
	i, j = 5, 8 // multiple assignment
	j, i = i, j // swap
	_, _ = i, j

}
