package main

import (
	"fmt"
)

func main() {


	var a string
	fmt.Printf("Input string: ")
	fmt.Scan(&a)
	runea := []rune(a)
	var b []rune
	for i:=len(runea)-1; i>=0; i-- {
		b = append(b, runea[i])
	}
	fmt.Println(string(b))
}