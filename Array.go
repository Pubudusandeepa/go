package main

import "fmt"

func main() {
	arr()
	addArray()
}

func arr() {
	var balance = [6]float32{1000.0, 2.0, 3.4, 7.0, 50.0, 12.0}

	for i := 0; i < 6; i++ {
		fmt.Print(balance[i])
	}
}

func addArray() {
	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Println("Array values", n[j])
	}
}
