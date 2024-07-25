package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	for {

		var userHeight float64 = 186
		var userKg float64 = 98

		fmt.Println("___Калькулятов индекса массы тела___")
		fmt.Print("Ввeдите свой рост в сантиметрах: ")
		fmt.Scan(&userHeight)

		fmt.Print("Введите свой вес: ")
		fmt.Scan(&userKg)

		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			fmt.Print(err)
			continue
		}

		outputResult(IMT)
		chooseRepeat := checkRepeatCalculation()
		if !chooseRepeat {
			break
		}
	}

}

func outputResult(imt float64) {
	result := fmt.Sprintf("\n___________________________________________\n\n\tВаш индекс массы тела: %.0f\n___________________________________________\n\n", math.Ceil(imt))

	fmt.Print(result)

	// if IMT < 16 {
	// 	fmt.Println("Девицит массы тела")
	// } else if IMT > 16 && IMT <= 18.5 {
	// 	fmt.Println("Недостаток веса")
	// } else if IMT > 18.5 && IMT <= 25 {
	// 	fmt.Println("Норма")
	// } else if IMT > 25 && IMT <= 30 {
	// 	fmt.Println("Избыточная масса")
	// } else if IMT > 30 && IMT <= 35 {
	// 	fmt.Println("1-я степень ожирения")
	// } else if IMT > 35 && IMT <= 40 {
	// 	fmt.Println("2-я степень ожирения")
	// } else if IMT > 40 {
	// 	fmt.Println("3-я степень ожирения")
	// }

	switch true {
	case imt < 16:
		fmt.Println("Девицит массы тела")
	case imt > 16 && imt <= 18.5:
		fmt.Println("Недостаток веса")
	case imt > 18.5 && imt <= 25:
		fmt.Println("Норма")
	case imt > 25 && imt <= 30:
		fmt.Println("Избыточная масса")
	case imt > 30 && imt <= 35:
		fmt.Println("1-я степень ожирения")
	case imt > 35 && imt <= 40:
		fmt.Println("2-я степень ожирения")
	default:
		fmt.Println("3-я степень ожирения")

	}
}

func calculateIMT(kg float64, height float64) (float64, error) {
	const IMTPower = 2
	if kg <= 0 || height <= 0 {
		var err = errors.New("NO_PARAMS_ERROR")
		return 0, err
	}
	var IMT = kg / math.Pow(height/100, IMTPower)

	return IMT, nil
}

func checkRepeatCalculation() bool {
	var result string
	fmt.Println("Повторить операцию? Введите Y")
	fmt.Scan(&result)
	if result == "y" || result == "Y" {
		return true
	}
	return false
}
