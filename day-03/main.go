package main

import (
	"fmt"
)

type Person struct {
		name string
		age int
}

type Rectangle struct {
		width , height int
}

	
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
	// prac4(y)
    // fmt.Println("value of the x is ", x)

	// p1 := Person{"John", 30}
	// p2 := &p1
	// fmt.Println("Before Value of the p1", p1.name)
	// p2.name = "Changed"
	// fmt.Println("After Value of the p1", p1.name)

	r := Rectangle{10, 5}
	fmt.Println("Area of the rectangle is ", r.area())
	r.scale(2)
	fmt.Println("Area of the rectangle after scaling is ", r.area())
	
}

// value receiver
func (r Rectangle) area() int {
	return r.width * r.height
}

func (r *Rectangle) scale(factor int) {
	r.width *= factor
	r.height *= factor
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