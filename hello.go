package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(max(11, 12))
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)
	greetings := []string{"Hello", "world!"}
	fmt.Println(strings.Join(greetings, " "))

}

func max(num1, num2 int) int {
	/* local variable declaration */

	result := 0

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}
