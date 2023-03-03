package main

import (
	"sync"
	"fmt"
)

func main() {
	wg := new(sync.WaitGroup)
	a := [5]int{2,4,6,8,10}
	sum := 0
	for _, elem := range a {
		wg.Add(1)
		go func(elem int, wg *sync.WaitGroup) {
			sum += elem * elem
			wg.Done()
		}(elem, wg)
	}
	wg.Wait()
	fmt.Println(sum)
}