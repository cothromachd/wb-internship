package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1")
		fmt.Println("2")
		runtime.Goexit()
		fmt.Println("3")
		fmt.Println("4")
	}()

	wg.Wait()
}