package main

import "fmt"

func main() {

	level := 7

	if level >= 10 {
		fmt.Println("Elite Adventurer")
	} else if level >= 5 {
		fmt.Println("Intermediate Adventurer")
	} else {
		fmt.Println("Beginner Adventurer")
	}
}
