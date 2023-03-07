package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for _, elem := range sequence {
		set[elem] = struct{}{}
	}

	

	fmt.Println(set)
}