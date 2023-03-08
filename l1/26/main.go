package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("input string: ")
	var s string
	fmt.Scan(&s)

	s = strings.ToLower(s)

	m := make(map[rune]int)
	for _, elem := range s {
		m[elem]++
	}

	flag := true
	for _, cnt := range m {
		if cnt > 1 {
			flag = false
		}
	}

	if flag {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}