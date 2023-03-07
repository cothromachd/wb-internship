package main

import (
	"fmt"
	"sync"
	//"sync"
	"time"
)

func main() {
	quit := time.After(5*time.Second)
	wg := new(sync.WaitGroup)
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-quit:
				fmt.Println("quitting")
				return
			default:
				fmt.Println("run in process...")
			}
		}
	}()

	wg.Wait()
}