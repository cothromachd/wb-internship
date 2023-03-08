package main

import "fmt"

func main() {
	a := []int{1,6,8,2,4,7,9,0,1,3,5,6}
	fmt.Println("slice:", a)
	
	fmt.Printf("input i to delete elem from slice: ")
	var n int
	fmt.Scan(&n)
	
	a = append(a[:n], a[n+1:]...)
	fmt.Println(a)
}