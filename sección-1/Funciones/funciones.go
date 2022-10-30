package main

import "fmt"

//como se define una funcion

func main() {
	//va a invocar una funcion
	//hello()
	//mostrarMensaje("gato")
	//addNumbers(1, 2, 3)
	//result := double(3)
	//fmt.Println(result)
	addResult, mulResult := addAndMultiply(2, 3)
	fmt.Println("suma de 2,3 es ", addResult)
	fmt.Println("multiplicacion de 2,3 es ", mulResult)
}

/*func hello() {
	//fmt.Println("Hola Mundo con go")
}

//trabajar con parametros

func mostrarMensaje(str string) {
	//fmt.Println(str) //va a mostrar en la consola el string que pasemos como parametro (es el str)
}

// otra forma
func addNumbers(x, y, z int) {
	a := x + y
	b := x + z
	c := y + z
	fmt.Println(a, b, c)
}
//retornar un valor
func double(x int) int {
	y := x * 2
	return y
}*/

//suma y multiplicacion

func addAndMultiply(i, j int) (int, int) {
	return i + j, i * j
}
