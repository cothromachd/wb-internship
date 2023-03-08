package main

import (
	"fmt"

	"github.com/cothromachd/wb-internship/l1/24/point"
)

func main() {
	pnt := point.NewPoint(1, 2)
	pnt_ := point.NewPoint(3,4)

	fmt.Println(point.FindDistance(pnt, pnt_))
}