package main

import "fmt"

func main() {
	//Array de numeros impares
	odd := [7]int{1, 3, 5, 7, 9, 11, 13}
	//usando la palabra clave range con bucle for para iterar sobre los elementos del array
	for i, item := range odd {
		//imprime el indice y los elementos
		fmt.Printf("odd[%d] = %d \n", i, item)
	}
}
