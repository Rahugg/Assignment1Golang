package html_read

import (
	"fmt"
)

func GiveRating() {
	fmt.Println()
	fmt.Println("Write the index of item which you want to change")
	var index int
	fmt.Scanln(&index)

	if index > len(dataProducts)-1 {
		fmt.Println("Sorry there is no such product")
	} else {
		fmt.Println("Write the new rating")
		var newRating float64
		fmt.Scanln(&newRating)
		dataProducts[index].Stars = (newRating + dataProducts[index].Stars*float64(dataProducts[index].NumberOfPeople)) / float64(dataProducts[index].NumberOfPeople)
		dataProducts[index].NumberOfPeople++;
		fmt.Println(dataProducts)
	}
}
