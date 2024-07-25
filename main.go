package main

import (
	"fmt"
	"math"
)

func main() {

	var userHeight float64 = 186
	var userKg float64 = 98

	fmt.Println("___Калькулятов индекса массы тела___")
	fmt.Print("Ввeдите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)

	outputResult(calculateIMT(userKg, userHeight))

}

func outputResult(imt float64) {
	result := fmt.Sprintf("\n___________________________________________\n\n\tВаш индекс массы тела: %.0f\n___________________________________________\n\n", math.Ceil(imt))

	fmt.Print(result)
}

func calculateIMT(kg float64, height float64) float64 {
	const IMTPower = 2
	var IMT = kg / math.Pow(height/100, IMTPower)

	return IMT
}
