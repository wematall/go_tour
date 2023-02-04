package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	// v := Vertex{3, 4} // original code
	v := &Vertex{3, 4} // fixed code

	a = f // a MyFloat implements Abser
	// a = &v // a *Vertex implements Abser // original code
	a = v // fixed code

	// In the folloing line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.

	a = v

	fmt.Println(a.Abs())
}
