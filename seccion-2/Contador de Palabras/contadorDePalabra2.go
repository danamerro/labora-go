package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func contadorPalabras(palabra string) map[string]int {
	mapPalabra := make(map[string]int)
	for _, palabra := range strings.Fields(palabra) {
		mapPalabra[palabra] += 1
	}
	return mapPalabra
}

func main() {
	wc.Test(contadorPalabras) 
}