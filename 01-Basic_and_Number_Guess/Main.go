// Jede Go Datei muss zu einem package gehören. In den ersten ist das normalerweise "main" (Standard).
package main

// "fmt" - Standard Input/Output Library
import ("fmt")

// Main Funktion welche immer als erstes ausgefürt wird.
func main()  {
	var age int
	fmt.Print("Bitte gebe dein Alter ein: ")
	fmt.Scan(&age)
	var name string // Variablen können mit dem Typ angegeben werden. Go kann es auch selbst erkennen.
	fmt.Print("Bitte gebe deinen Namen ein: ")
	fmt.Scan(&name)
	// fmt.Println("Hallo " + name)
	// fmt.Printf("Hallo, du bist %v Jahre alt!", age) //Printf wird benötigt um Tesxt mit einer Zahl (var age) auszugeben.
	// fmt.Printf("Hallo " + name + " du bist %v Jahre alt!", age) // Meine Variante
	fmt.Printf("Hallo %v du bist %v Jahre alt!\n", name, age) // Verbessert im Video

	if age < 18 {
		var missing = 18 - age
		var ages = "Jahre"
		if missing == 1 {
			ages = "Jahr"
		}
		fmt.Printf("Du bist NICHT volljährig! Noch %v %v bis dahin.", missing, ages)
	} else {
		fmt.Println("Du bist volljährig!")
	}
}