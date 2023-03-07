package main

import (
	"fmt"
	"sync"
)

type counter struct {
	cnt int
	sync.Mutex
}

func (c *counter) inc() {
	defer c.Unlock()
	c.Lock()
	c.cnt++
}


func main() {
	cnter := counter{cnt: 0}
	wg := new(sync.WaitGroup)

	for i:=0; i<1000;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnter.inc()
		}()
	}
	wg.Wait()
	fmt.Println(cnter.cnt)
}