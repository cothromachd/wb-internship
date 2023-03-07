package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := time.After(8 * time.Second)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			v, ok := <-ch
			if ok != false {
				fmt.Println(v)
			} else {
				break
			}
		}
	}()

	breaker := false
	for {
		select {
		case <-timeout:
			close(ch)
			breaker = true
			break
		default:
			ch <- 1
		}

		if breaker == true {
			break
		}
	}
	
	wg.Wait()
}