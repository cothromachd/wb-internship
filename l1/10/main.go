package main

import (
	"fmt"
)

func main() {
	a := []float64{-30.1, -30, -29.9, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 5, 0.1, 9.9, -0.1, -9.9}

	m := make(map[int][]float64)

	for _, elem := range a {
		elem_ := int(elem)
		m[elem_/10] = append(m[elem_/10], elem)
	}

	fmt.Println(m)
}