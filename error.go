package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := sqrt(4)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}

	return math.Sqrt(value), nil
}
