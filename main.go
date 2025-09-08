package main

import (
	"fmt"
	"math"
)

func main() {

	userHeight, userKg := getUserInput()

	var IMT = calculateIMT(userKg, userHeight)

	outputResult(IMT)

}

func outputResult(imt float64) {
	result := fmt.Sprintf("Index body mass: %.2f", imt)
	fmt.Println(result)
}

func calculateIMT(weight, height float64) float64 {
	const IMTPower = 2

	return weight / math.Pow(height, IMTPower)
}

func getUserInput() (float64, float64) {
	var height float64
	var weight float64

	fmt.Println("Enter you height:")
	fmt.Scan(&height)
	fmt.Println("Enter you weihgt:")
	fmt.Scan(&weight)

	return height, weight
}
