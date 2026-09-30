package main

import "fmt";

func main(){
	// defer fmt.Println("world")
	// defer fmt.Println("world 1")
	// fmt.Println("Hello")
	// mydefer()
	 defer handlePanic()

    panic("something went wrong")
}

func mydefer(){
	for i:=0 ; i<=5; i++ {
		defer fmt.Println(i)
	}
}

func handlePanic() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}
