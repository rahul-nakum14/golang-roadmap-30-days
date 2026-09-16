package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	user := Person{"John", 30}
	// printUser(user)
	// createdUser := createUser("Alice", 25)
	// fmt.Println("Created user:", createdUser.name, "Age:", createdUser.age)
	// user.newPrintUser();
	// user.changeName("Rahul")
	// user.changeNamePointer("Rahul")
	// fmt.Println("After changeNamePointer method, name is:", user.name)
	result := user.getInfo();
	fmt.Println("Retrieved info:", result);
}

//passes the struct by value.
func printUser(user Person){
	fmt.Println("Hey", user.name, "you are", user.age, "years old.")
}

//Function that returns a struct
func createUser(name string, age int) Person {
	return Person{
		name: name,
		age:  age,
	}
}

// Introduce a method with a value receiver
func(u Person) newPrintUser(){
	fmt.Println("Hey", u.name, "you are", u.age, "years old.")
}

// Introduce a method but with the parameter
func(u Person) changeName(newName string){
	u.name = newName
	fmt.Println("Inside changeName method, name changed to:", u.name)
}

// Introduce a method with a pointer receiver
func(u *Person) changeNamePointer(newName string){
	u.name = newName
	fmt.Println("Inside changeNamePointer method, name changed to:", u.name)
}

func (u *Person) getInfo() string {
	return u.name
}