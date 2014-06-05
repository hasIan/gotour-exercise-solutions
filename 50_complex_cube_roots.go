package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) complex128 {
    z := x
    for i := 0; i < 10; i++ {
        z -= (cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2))
    }

    return z
}

func main() {
    fmt.Println(Cbrt(2))
    fmt.Println(Cbrt(complex128(3 + 4i)))
}
