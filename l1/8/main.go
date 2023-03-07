package main

import (
	"fmt"
	"math"
)

func main() {
	// 100000 / 101000
	// 10000
	// 10000
	var num int64 = math.MaxInt64
	fmt.Printf("number: %d\n", num)
	fmt.Printf("Input bit and index: ")
	var bit uint8
	var index uint8
	fmt.Scan(&bit, &index)
	if bit == 1 { // 
				  // 100000 | 100100 = 100100
		num = num | (1<<index)
	} else if bit == 0 { // 100100 & (100100 - 000100) = 100000
		num = num & (num - (1 << index))
	} else {
		fmt.Println("error: wrong bit")
	}
	fmt.Println("output number:", num)
}

  