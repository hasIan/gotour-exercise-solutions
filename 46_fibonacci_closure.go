package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 0
	return func() int {
		if y == 0 {
			y = 1
		} else {
			x, y = y, x+y
		}
		return y
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Println(f())
	}
}
