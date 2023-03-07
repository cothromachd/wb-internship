package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	ch_ := make(chan int)
	wg := new(sync.WaitGroup)
	a := []int{1,2,3,4,5,6,7,8,9,10}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			v, ok := <-ch
			if ok != false {
				ch_ <- v * 2
			} else {
				close(ch_)
				break
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			v, ok := <-ch_
			if ok != false {
				fmt.Println(v)
			} else {
				break
			}
		}
	}()
	
	for _, elem := range a {
		ch <- elem
	}
	close(ch)
	
	wg.Wait()
}