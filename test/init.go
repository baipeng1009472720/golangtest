package test

import (
	"math"
	"fmt"
)

var Pi float64

func Init() {
	Pi = 4 * math.Atan(1) // init() function computes Pi
	fmt.Println(Pi)
}
