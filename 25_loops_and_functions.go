package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const threshold = 0.001

	z, zPrime := x, x
	for i := 0; i < 100; i++ {
		zPrime = z
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
		if math.Abs(z-zPrime) < threshold {
			fmt.Printf("i = %d\n", i)
			break
		}
	}

	return z
}

func main() {
	const operand = 2

	fmt.Println(Sqrt(operand))
	fmt.Println(math.Sqrt(operand))
}
