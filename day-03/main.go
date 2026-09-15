package main

import (
	"fmt"
)

func main(){
	x := 10
	y := &x

	// fmt.Println("value of the x is ", x)
	// fmt.Println("value of the y is ", y)
	// fmt.Println("value of the *y is ", *y)

	*y = 20
	// fmt.Println("value of the x is ", x)
	// fmt.Println("value of the y is ", y)
	// fmt.Println("value of the *y is ", *y)
	// prac1()
	// prac2()
	// prac3()
	prac4(y)
    fmt.Println("value of the x is ", x)
}


func prac1() {
    x := 50
    p := &x

	*p = 100
    fmt.Println(x)
}

func prac2() {
    x := 10

    p1 := &x
    p2 := &x // p1,p2 == 10

    *p1 = 20 
    *p2 = 30

    fmt.Println(x) // 30
    fmt.Println(*p1) // 30
    fmt.Println(*p2) // 30
}

func prac3() {
	x := 10
	p := &x
	pp := &p

	fmt.Println(x)
	fmt.Println(*p)
	fmt.Println(**pp)
}

func prac4(x *int) {
	*x = 20
}