package main

import (
	"fmt"
)

func worker(c <-chan int) {
	for {
		fmt.Println(<-c)
	}

}


func main() {
	c := make(chan int)

	var n int
	fmt.Printf("Type number of workers: ")
	fmt.Scan(&n)

	for i:=0; i<n; i++ {
		go func() {
			for {
				fmt.Println(<-c)
			}
		}()
	}

	for {
		c <- 1
	}
}