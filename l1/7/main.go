package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	for i:=0; i<1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}