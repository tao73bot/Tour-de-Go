package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // v is called the "receiver"
	return math.Abs(v.X + v.Y)
}

func abs(v Vertex) float64 { // function with Vertex as argument
	return math.Abs(v.X + v.Y)
}

func (v *Vertex) Scale(f float64) { // pointer receiver
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{-3, -4}
	v.Scale(10)
	fmt.Println(v.Abs()) // method call
	fmt.Println(abs(v))  // function call

	// You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
	// You can declare a method with a receiver whose type is defined in another package, but not in the same package.
	// You can declare a method with a receiver whose type is defined in the same package, but not in another package.

}
