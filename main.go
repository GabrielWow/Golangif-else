package main

import "fmt"

func main() {
	// day of a month [1, 31]

	/*commentaire
	sur plusieurs
	lignes.
	*/

	day := 18

	if day < 15 {

		fmt.Printf("We're in the first half of the month (day=%d)\n", day)
	} else if day == 18 {
		fmt.Printf("We're in the special day (day=%d)\n", day)
	} else {
		fmt.Printf("We're in the second half of the month (day=%d)\n", day)
	}

	//booléen

	year, month, day := 2009, 11, 10
	fmt.Printf("Date=%d/%d/%d\n", day, month, year)

	// Test sur les dates
	if year == 2009 && month == 11 && day == 10 {
		fmt.Println("This is the first relase of GO!")
	} else if year == 2009 || month == 11 || day == 10 {
		fmt.Println("At least one part is right... :/")
	} else {
		fmt.Println("Just another day...")
	}

	// Capacité de déclarer des variables a l'entérieur d'un if

	if count := 12; count > 10 {
		fmt.Printf("We have enough count. got= %d\n", count)
	} else {
		fmt.Printf("Not enough. got=%d\n", count)
	}

}
