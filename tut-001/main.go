package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strconv"
	"strings"
	// "sort"
)

func main(){
	// timing()
	// userINput()
	// converSion()
	// pointer()
	// arrays()
	// slices()
	maps()
}

func arrays(){
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	fmt.Println("array is ", arr)
	fmt.Println("array length is ", len(arr))

	var data = [5]string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("data is ", data)
	fmt.Println("data length is ", len(data))
}

func pointer(){
	var num int = 42
	var ptr *int = &num
	fmt.Println("num is ", num)
	fmt.Println("ptr is ", ptr)
	fmt.Println("value at ptr is ", *ptr) 
}

func converSion(){
	fmt.Println("rate between 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	rate, _ := reader.ReadString('\n')
	fmt.Println("your rate is ", rate)

	// 1 .st  way 
	// numRating, err := strconv.ParseFloat(strings.TrimSpace(rate), 64)
	// if err != nil {
	// 	fmt.Println("Error parsing rate:", err)
	// 	return
	// }
	// fmt.Println("your rate + 1 is ", numRating + 1)

	// 2 .nd way
	numRating, err := strconv.Atoi(strings.TrimSpace(rate))
	if err != nil {
		fmt.Println("Error parsing rate:", err)
		return
	}
	fmt.Println("the rating is the numRating", numRating+2)

	// ParseFloat is used for decimal numbers like 3.4, while Atoi is for whole numbers like 3.

}

func userINput(){
	//`fmt.Scanln()` reads input only up to whitespace,
	// while `bufio.Reader.ReadString('\n')` reads the entire line, 
	// including spaces, until Enter is pressed.
	var name string
	fmt.Println("Enter your name")
	fmt.Scanln(&name)
	fmt.Println("your name is ", name)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your address")
	address, _ := reader.ReadString('\n')
	fmt.Println("your address is ", address)
}

func timing(){
	parseTime := time.Now();
	fmt.Println("parsetime is ->", parseTime)
	fmt.Println("time format is ->", parseTime.Format("2006-01-02 15:04:05 Monday"))

	createdDate := time.Date(2020,time.May,14,11,34,0,0, time.UTC)
	fmt.Println("created datee" , createdDate.Format("2006-01-02 15:04:05 Monday"))
}

func slices() {
	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println("slice is:", numbers)
	fmt.Println("length:", len(numbers))
	fmt.Println("capacity:", cap(numbers))

	// adding element
	numbers = append(numbers, 60)
	fmt.Println("after append:", numbers)

	//adds multiple elements
	numbers = append(numbers, 70, 80)
	fmt.Println("after multiple append:", numbers)

	// access elements
	fmt.Println("first element:", numbers[0])
	fmt.Println("last element:", numbers[len(numbers)-1])

	// Slicing data
	fmt.Println("first three:", numbers[:3])
	fmt.Println("from index 2:", numbers[2:])

	// Modify element
	numbers[0] = 100
	fmt.Println("after modification:", numbers)

	data := make([]int, 3)
	data[0] = 23
	data[1] = 43
	data[2] = 32
	// sort.Ints(data)
	// fmt.Print("Are the elements sorted? ", sort.IntsAreSorted(data))
	fmt.Println("data slice:", data)

	var index int = 1
	// Remove element at index 1
	data = append(data[:index], data[index+1:]...)
	fmt.Println("after removing element at index 1:", data)
}


func maps() {
	// Create a map to store student names and their ages
	students := make(map[string]int)
	students["Alice"] = 20
	students["Bob"] = 22
	students["Charlie"] = 19

	fmt.Println("students -->", students)
	fmt.Println("Alice's age:", students["Alice"])

	// Update a student's age
	students["Bob"] = 23
	fmt.Println("Updated students -->", students)
}