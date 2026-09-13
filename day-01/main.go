// There are 3 ways to declare a variable in golang

// 1 . var variable_name data_type
// 2 . var variable_name
// 3 . variable_name

package main

import "fmt"

const sal = "dfs";

func main(){

	// 1 st method
	var age int;
	age = 18

	fmt.Println("Age is --->", age)

	var isbool bool

	fmt.Println("default val is the ", isbool) //"false"

	var name string
	name = "rahul"
	fmt.Println("name --- > ", name) // ""

	var num int
	fmt.Println("num is --->",num) // 0

	var floatres float32
	fmt.Println("default float is", floatres)

	// 2 nd method - go automatic intefer type

	var result = "pass"
	fmt.Println("result --> ", result)

	var floatResult = 12.4
	fmt.Println("float result -->", floatResult)

	// 3 short variable name declarations

	ageDetails := 13
	fmt.Println("ageDetails ->", ageDetails) 

	floatData := 33.32
	fmt.Printf("flaotData %T->", floatData) // default take float 64

	fmt.Println("\n")
	 sal:="asdasd"
	//const variable
	fmt.Println("sal is ->",sal)
	// print vs println

	// print always result in a single line output while println result in a new seprate line

	fmt.Print("hello ")
	fmt.Print("world") // hello world

	fmt.Print("\n")
	
	fmt.Println("hello ")
	fmt.Println("world") // hello world

	// printf is used to format string 
	fmt.Printf("Name: %s, Age: %d , data type of name: %T", name, age, name)
	
}

