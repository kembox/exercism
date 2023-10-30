package darts

import (
	"math"
)

func Score(x, y float64) int {
	d := math.Sqrt(x*x + y*y)
	if d > 10 {
		return 0
	} else if d > 5 {
		return 1
	} else if d > 1 {
		return 5
	} else {
		return 10
	}
}