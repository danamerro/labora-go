package main

import "fmt"

func main() {
	//Boolean
	var b1 bool = true //declaracion tipada con valor inicializado
	var b2 = true      // no tipada pero con valor inicializado
	var b3 bool        //tipada pero sin valor inicializado
	b4 := true         //no tipada pero con valor inicializado

	fmt.Println("parte de los booleans")
	fmt.Println(b1) //devuelve true
	fmt.Println(b2) //devuelve true
	fmt.Println(b3) //devuelve false
	fmt.Println(b4) //devuelve true

	fmt.Println("--------------------")
	//Integer (int)
	var numeroDeLunasInt int
	numeroDeLunasInt = 21
	fmt.Println("parte de los int")
	//devuelve Data=21 , Type = int
	fmt.Printf("Data = %v , Type = %T", numeroDeLunasInt, numeroDeLunasInt)

	//Byte
	//se utiliza para representar un entero de 8 bits sin signo. El tipo de dato de byte por defecto es 0
	var numeroDeLunas byte
	//a la variable tipo byte se asigna el num 1.
	numeroDeLunas = 1
	fmt.Println()
	fmt.Println("--------------------")
	fmt.Println("Parte de los Bytes")
	fmt.Printf("Data = %v, Type = %T", numeroDeLunas, numeroDeLunas)

	//bytes simples
	var a1 byte = 97
	var a2 byte = 98
	var a3 byte = 99

	//imprime contenido de la variable
	fmt.Println()
	fmt.Println("************Imprime el contenido de la variable******")
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	//imprime la representacion de caracteres del byte (ASCII)
	fmt.Println("************Imprime la representacion de caracteres******")
	fmt.Printf("%c\n", a1)
	fmt.Printf("%c\n", a2)
	fmt.Printf("%c\n", a3)

	//Float
	//existen dos float32 y float64 bits
	fmt.Println("--------------------")
	fmt.Println("Float")
	var peso float32
	peso = 17.23
	fmt.Printf("Data = %v, Type = %T", peso, peso)
	//String
	fmt.Println()
	fmt.Println("--------------------")
	fmt.Println("String")
	var message string
	message = "Hola..."
	fmt.Printf("Data = %v , Type = %T", message, message)
}
