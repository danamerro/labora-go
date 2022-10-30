package main

import "fmt"

func main() {
	grade := 70

	if grade >= 65 {
		fmt.Println("Passing grade")
	}

	//for
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}
	for i := 0; i < len(sharks); i++ {
		fmt.Println(sharks[i])
	}

}
