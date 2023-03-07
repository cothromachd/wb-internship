package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("finish")
				return
			}
			fmt.Println(v)
		}
	}()
   
	ch <- struct{}{}
	close(ch)
	
	time.Sleep(time.Second)
   }