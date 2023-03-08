package main

import (
	"fmt"
	"time"
)

func Sleep(x int) {
	timer := time.After(time.Second * 5)

	<-timer
}

func main() {
	Sleep(5)
	fmt.Println(5)
}