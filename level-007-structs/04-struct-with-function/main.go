package main

import "fmt"

type Character struct {
	Name  string
	Level int
	Title string
}

func displayCharacter(character Character) {
	fmt.Println("Name:", character.Name)
	fmt.Println("Level:", character.Level)
	fmt.Println("Title:", character.Title)
}

func main() {

	character := Character{
		Name:  "Charles",
		Level: 7,
		Title: "Collection Handler",
	}

	displayCharacter(character)
}
