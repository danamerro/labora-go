package main

import "fmt"

type PlanetayLunas struct {
	Name string
	Moon int
}

func main() {
	var planeta1 = PlanetayLunas{Name: "Mercurio", Moon: 0}
	var planeta2 = PlanetayLunas{Name: "Venus", Moon: 0}
	var planeta3 = PlanetayLunas{Name: "Tierra", Moon: 1}
	var planeta4 = PlanetayLunas{Name: "Marte", Moon: 2}
	var planeta5 = PlanetayLunas{Name: "Jupiter", Moon: 63}
	var planeta6 = PlanetayLunas{Name: "Saturno", Moon: 62}
	var planeta7 = PlanetayLunas{Name: "Urano", Moon: 27}
	var planeta8 = PlanetayLunas{Name: "Neptuno", Moon: 13}

	lista := []PlanetayLunas{planeta1, planeta2, planeta3, planeta4, planeta5, planeta6, planeta7, planeta8}

	for i := 0; i < len(lista); i++ {
		if lista[i].Moon > 0 {
			fmt.Println("El planeta", lista[i].Name, "tiene", lista[i].Moon, " lunas")
		} else {
			fmt.Println("El planeta", lista[i].Name, " "+"no tiene lunas")
		}
	}
}
