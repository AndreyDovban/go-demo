package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64

	fmt.Println("Enter you height:")
	fmt.Scan(&userHeight)
	fmt.Println("Enter you weihgt:")
	fmt.Scan(&userKg)

	var IMT = userKg / math.Pow(userHeight, IMTPower)

	fmt.Printf("Index body mass: %.2f", IMT)
}
