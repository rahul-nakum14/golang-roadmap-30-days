package main

import (
	"fmt"
)

func main() {

	// var arrayDemo[4]int = [4]int{0, 0, 0, 0}
	// var arrayDemo = [4]int{0, 0, 0, 0}
	// var arrayDemo[4]int

	// arrayDemo[0] = 10
	// arrayDemo[1] = 20
	// arrayDemo[2] = 30
	// arrayDemo[3] = 40

	// fmt.Println(arrayDemo)

	// var arrSlice = []int{10, 20, 30, 40, 50}
	// fmt.Println(arrSlice)
	// arrSlice = append(arrSlice, 60)
	// fmt.Println(arrSlice)
	
	// arrSlice = append(arrSlice[1:3], 70, 80, 90)
	// fmt.Println(arrSlice)

	arrNew := make([]int, 5, 10)
	fmt.Println(arrNew)
	fmt.Println(len(arrNew))
	fmt.Println(cap(arrNew))
}