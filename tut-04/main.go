package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int `json:"age"`
	Salary float64 `json:"salary"`
	Password string `json:"-"` // This field will be ignored during JSON marshaling
	Books []string `json:"books,omitempty"` // This field will be omitted if it's empty
}

func main() {

	// personDetails := []Person{
	// 	{Name: "John", Age: 30, Salary: 50000.0, Password: "john123", Books: []string{"Book1", "Book2"}},
	// 	{Name: "Alice", Age: 25, Salary: 60000.0, Password: "alice123", Books: []string{"Book3", "Book4"}},	
	// 	{Name: "Bob", Age: 35, Salary: 70000.0, Password: "bob123", Books:nil}, // as here Books field is empty, so it will be omitted
	// }
	// result , _ := json.MarshalIndent(personDetails, "", "  ")
	// fmt.Println("JSON representation:", string(result))
	decodeJSON()
}

func decodeJSON() {
	jsonData := `[
		{"name": "John", "age": 30, "salary": 50000.0, "books": ["Book1", "Book2"]},
		{"name": "Alice", "age": 25, "salary": 60000.0, "books": ["Book3", "Book4"]},
		{"name": "Bob", "age": 35, "salary": 70000.0}
	]`
	var persons []Person

	checkVar := json.Valid([]byte(jsonData))
	if checkVar {
		fmt.Println("jSON is valid")
		err := json.Unmarshal([]byte(jsonData), &persons)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println("Decoded persons:", persons)
		} else {
				fmt.Println("JSON is invalid")
		}

		var mapData []map[string]interface{}
		err := json.Unmarshal([]byte(jsonData), &mapData)
		if err != nil {
			fmt.Println("Error decoding JSON into map:", err)
			return
		}
		fmt.Printf("Decoded map data: %#+v\n", mapData)

		for key, value := range mapData {
			fmt.Printf("key is %d and value is %v and the type is %T\n", key, value, value)
		}
}