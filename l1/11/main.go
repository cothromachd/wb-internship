package main

import "fmt"

func main() {
	a := []int{1,2,3,4,5,6,7,8,9}
	b := []int{5,6,7,8,9,10,11,12,13}
	c := []int{}

	m := make(map[int]int)

	for _, elem := range a {
		m[elem]++
	}

	for _, elem := range b {
		m[elem]++
	}

	for key, value := range m {
		if value != 1 {
			c = append(c, key)
		}
	}
	
	fmt.Println(c)
}