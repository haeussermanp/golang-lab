package main

import (
	"fmt"
	"math/rand" // Contains functions for generating pseudo-random numbers.
)

const defaultMin = 1
const defaultMax = 100

func main() {
	var current_guess int
	var difference string

	start_number := rand.Intn(defaultMax - defaultMin) + defaultMin

	for current_guess != start_number {
		fmt.Println("Bitte gebe deine aktuelle Schätzung ein: ")
		fmt.Scan(&current_guess)

		if current_guess != start_number {
			if current_guess < start_number {
				difference = "NIEDRIG"
			} else if current_guess > start_number {
				difference = "HOCH"
			}
			fmt.Printf("Leider war die Schätzung nicht richtig. Deine Zahl war leider zu %v. \n", difference)
		}
	}

	if current_guess == start_number {
		fmt.Printf("Deine Schätzung %v stimmt mit der Start Nummer %v überein. Herzlichen Glückwunsch!", current_guess, start_number)
	}
}