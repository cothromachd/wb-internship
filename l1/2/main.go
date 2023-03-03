package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	a := [5]int{2,4,6,8,10}
	for _, elem := range a {
		wg.Add(1)
		go func(elem int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(elem * elem)
		}(elem, wg)
	}
	wg.Wait()
}