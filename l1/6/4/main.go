package main

import (
	"fmt"
)

func main() {
	
	set := make(map[string]bool) // New empty set
	set["Foo"] = true            // Add
	for k := range set {         // Loop
		fmt.Println(k)
	}
	delete(set, "Foo")    // Delete
	size := len(set)      // Size
	exists := set["Foo"]
	fmt.Println(size, exists)
}