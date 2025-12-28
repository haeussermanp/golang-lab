package main

import ("fmt")

func main() {
	var start_number int
	var current_guess int

	fmt.Println("Bitte gebe die Nummer ein die erraten werden soll: ")
	fmt.Scan(&start_number)

	for current_guess != start_number {
		fmt.Println("Bitte gebe deine aktuelle Schätzung ein: ")
		fmt.Scan(&current_guess)

		if current_guess != start_number {
			fmt.Println("Leider war die Schätzung nicht richtig. ")
		}
	}

	if current_guess == start_number {
		fmt.Printf("Deine Schätzung %v stimmt mit der Start Nummer %v überein. Herzlichen Glückwunsch!", current_guess, start_number)
	}
}