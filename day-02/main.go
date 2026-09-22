package main

import (
	"fmt"
	"time"
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

	// range statement

	nums:= []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		switch value {
		case 1:
			fmt.Println("Skipping the value:", value)
		default:
			fmt.Println("Index:", index, "Value:", value)
		}
	}

	// Function call
	// var n int;
	// fmt.Print("Enter a number to calculate factorial: ")
	// fmt.Scan(&n)
	// result := factorial(n)
	// fmt.Printf("Factorial of %d is %d\n", n, result)

	parseTime := time.Now()
	fmt.Println("Current time is:", parseTime.Format("2006-01-02 15:04:05 Monday"))

	createDate := time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Created date is:", createDate.Format("2006-01-02 15:04:05 Monday"))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
