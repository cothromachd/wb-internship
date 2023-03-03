package main

import (
	"fmt"
	"time"
)

func worker(c <-chan int) {
	for {
		fmt.Println(<-c)
	}

}


func main() {
	c := make(chan int)
	ticker := time.NewTicker(1 * time.Second)

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

	for range ticker.C {
		c <- 1
	}
}