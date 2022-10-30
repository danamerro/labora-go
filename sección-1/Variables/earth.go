package main

import "fmt"

func main() {
	//manera de declarar una variable
	planetName := "Earth"
	numberOfMoons := 1
	fmt.Println(planetName, numberOfMoons)
	//Declaracion de multiples variables
	var i, j, k = 1, 2, 3

	var (
		name       = "Jhon Doe"
		occupation = "gardener"
	)
	fmt.Println(i, j, k)
	fmt.Println("------")
	fmt.Printf("%s is a %s", name, occupation)
	//variable
	var age int = 34
	//constante
	const WIDTH = 100
	//se modifica la edad
	age = 35
	age = 36
	fmt.Println(age, WIDTH)

}
