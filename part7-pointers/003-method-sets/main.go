package main

import (
	"fmt"
	"math"
)

// only works with values that are pointers
// receivers     values
// (t T)         T and *T
// (t *T)        *T
type circle struct {
	radius float64
}
type shape interface {
	area() float64
}
func (c *circle) area() float64  {
	return math.Pi * c.radius * c.radius
}

func info(s shape){
	fmt.Println(s.area())
}

func main()  {
	a := circle{
		radius: 5,
	}
	info(&a)
	fmt.Println(a.area())
}