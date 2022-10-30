package main

import "fmt"

// Struct
type Videogame struct {
	Title string
	Year  int
}

func main() {
	var myVideoGame = Videogame{Title: "Nier", Year: 2017}
	fmt.Println(myVideoGame.Title, "Year ", myVideoGame.Year)

}
