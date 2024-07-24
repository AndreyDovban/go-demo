package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight = 1.57
	var userKg float64 = 54
	var IMT = userKg / math.Pow(userHeight, IMTPower)

	fmt.Print(IMT)
}
