package main

import "fmt"

type Character struct {
	Name  string
	Level int
	Title string
}

func main() {
	character := Character{
		Name:  "Charles",
		Level: 7,
		Title: "Collection Handler",
	}

	character.Level = 8
	character.Title = "Model Architect"

	fmt.Println(character)
}
