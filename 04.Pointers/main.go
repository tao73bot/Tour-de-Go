package main

import "fmt"

type Point struct {
	X int
	Y int
}

var(
	v1 = Point{1, 2}  // has type Point
	v2 = Point{X: 1}  // Y:0 is implicit
	v3 = Point{}      // X:0 and Y:0
	p  = &Point{1, 2} // has type *Point
)

func main() {
	// The type *T is a pointer to a T value. Its zero value is nil.
	// The & operator generates a pointer to its operand.

	i, j := 12, 23

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p * 2    // multiply j through the pointer
	fmt.Println(j) // see the new value of j

	// Structs
	p1 := Point{1, 2}
	p1.X = 1e9
	fmt.Println(p1)
	// Struct pointers
	pp := &p1
	pp.Y = 1e3 // shorthand for (*pp).Y = 1e3
	fmt.Println(p1)

	// Struct literals
	fmt.Println(v1, p, v2, v3)
}
