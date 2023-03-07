package main

import(
	"sort"
	"fmt"
)

func main() {
	a := []int{9, 6, 1, 2, 4, 7, 0, 13, 19, 20, 3}
	sort.Ints(a)

	fmt.Println(a)
}