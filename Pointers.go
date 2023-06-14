package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("Address of a variable: %x\n", &a)
	pointers()
	pointerstopointers()
}

func pointers() {
	var age int = 20
	var val *int
	val = &age

	fmt.Println("address of val ", val)
	fmt.Println("value of val ", *val)
}

func pointerstopointers() {
	var grade int = 10
	var point *int
	var ptop **int

	point = &grade
	ptop = &point

	fmt.Println(*point)
	fmt.Println(**ptop)
}
