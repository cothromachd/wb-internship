package point

import (
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(xp, yp float64) *Point {
	return &Point{xp, yp}
}

func FindDistance(pp *Point, pp_ *Point) float64 {
	return math.Sqrt(math.Pow(pp_.y-pp.y, 2) + math.Pow(pp_.x-pp.x, 2))
}

func main() {
	
}