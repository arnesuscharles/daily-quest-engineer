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

	fmt.Println(character)
}
