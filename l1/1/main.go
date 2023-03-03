package main

import "fmt"

type Human struct {
	Name string
	Surname string
	Age int
	Height int
	Weight int
}

func(h *Human) About() {
	fmt.Println(h.Name, h.Surname, h.Age, h.Height, h.Weight)
}

type Action struct {
	Human
}


func main() {
	a := Action{Human: 
		Human{
		Name: "Khalid", Surname: "Dzhamaldaev", Age: 19, Height: 182, Weight: 72,
	}}
	a.About()
}