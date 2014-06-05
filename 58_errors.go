package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(f float64) (float64, error) {

    if f < 0 {
        return 0, ErrNegativeSqrt(f)
    }

    const threshold = 0.001

    z, zPrime := f, f
    for i := 0; i < 100; i++ {
        zPrime = z
        z = z - ((math.Pow(z, 2) - f) / (2*z))
        if math.Abs(z - zPrime) < threshold {
            break
        }
    }

    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
