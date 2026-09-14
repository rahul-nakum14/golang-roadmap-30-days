package main

import (
	"fmt"
)

func main() {

	// If else statement

	age := 20

	if age >= 18 {
		fmt.Println("You are an adult.")
	}else{
		fmt.Println("You are a minor.")
	}


	// Swith case statement

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday.")
	case "Tuesday":
		fmt.Println("It's Tuesday.")
	default:
		fmt.Println("It's another day.")
	}

	for i := 1; i <= 5; i++ {
		switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		default:
			fmt.Println("Other number")
		} 
	}
}
