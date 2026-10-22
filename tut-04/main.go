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

	personDetails := []Person{
		{Name: "John", Age: 30, Salary: 50000.0, Password: "john123", Books: []string{"Book1", "Book2"}},
		{Name: "Alice", Age: 25, Salary: 60000.0, Password: "alice123", Books: []string{"Book3", "Book4"}},	
		{Name: "Bob", Age: 35, Salary: 70000.0, Password: "bob123", Books:nil}, // as here Books field is empty, so it will be omitted
	}

	result , _ := json.MarshalIndent(personDetails, "", "  ")
	fmt.Println("JSON representation:", string(result))

}