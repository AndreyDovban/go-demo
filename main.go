package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recover ", r)
		}

	}()

	for {
		userHeight, userKg := getUserInput()

		var IMT, err = calculateIMT(userKg, userHeight)
		if err != nil {
			// fmt.Println(err.Error())
			// continue
			panic("Не заданы параметры для расчёта")
		}

		outputResult(IMT)

		if !questionForContinue() {
			break
		}

	}

}

func questionForContinue() bool {
	var response string
	fmt.Println("Хотите продлжить? y/n")
	fmt.Scan(&response)

	if strings.ToLower(response) == "y" || strings.ToLower(response) == "yes" {
		return true
	}
	return false
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Index body mass: %.2f", imt)
	fmt.Println(result)
	switch true {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас норма")
	case imt < 30:
		fmt.Println("У вас избыточная масса тела")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(weight, height float64) (float64, error) {
	const IMTPower = 2

	if weight <= 0 || height <= 0 {
		return 0, errors.New("ПРОГРАММНАЯ ОШИБКА")
	}

	result := weight / math.Pow(height/100, IMTPower)

	return result, nil
}

func getUserInput() (float64, float64) {
	var height float64
	var weight float64

	fmt.Println("Enter you height:")
	fmt.Scanln(&height)
	fmt.Println("Enter you weihgt:")
	fmt.Scan(&weight)

	return height, weight
}
