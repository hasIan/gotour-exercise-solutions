package main

import "code.google.com/p/go-tour/pic"
import "math"

var calc = avg

func Pic(dx, dy int) [][]uint8 {
    y := make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        x := make([]uint8, dx)
        for j := 0; j < dx; j++ {
            x[j] = uint8(calc(i, j))
        }
        y[i] = x
    }

    return y
}

func avg(v1, v2 int) float32 {
    return float32((v1 + v2) / 2)
}

func mult(v1, v2 int) int {
    return v1 * v2
}

func pow(v1, v2 int) float64 {
    return math.Pow(float64(v1), float64(v2))
}

func main() {
    pic.Show(Pic)
}
