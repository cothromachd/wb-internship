package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var a_ []byte
	fmt.Printf("Input string: ")
	a_, _, _ = bufio.NewReader(os.Stdin).ReadLine()
	a := string(a_)
	as := strings.Split(a, " ")
	
	var b []string
	for i:=len(as)-1;i>=0;i-- {
		b = append(b, as[i])

	}

	bj := strings.Join(b, " ")
	fmt.Println(bj)
}