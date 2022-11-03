package main

import "fmt"

func main() {
	var diasSemana map[int]string
	diasSemana = make(map[int]string)
	diasSemana[1] = "Domingo"
	diasSemana[2] = "Lunes"
	diasSemana[3] = "Martes"
	diasSemana[4] = "Miércoles"
	diasSemana[5] = "Jueves"
	diasSemana[6] = "Viernes"
	diasSemana[7] = "Sábado"

	fmt.Println(diasSemana)
}
