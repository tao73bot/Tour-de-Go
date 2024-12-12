package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var i, j int = 1, 2

// x := 3 // this will not work outside of a function because it is a short variable declaration

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14 // Constants are declared like variables, but with the const keyword.
// const world string := "世界" // Constants cannot be declared using the := syntax

// Numeric Constants

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	var c, python, java = true, false, "no!"
	k := 3 // short variable declaration
	fmt.Printf("%T %T %T %T %T %T\n", i, j, c, python, java, k)
	fmt.Println(i, j, c, python, java, k)

	fmt.Printf("%T %T %T\n", ToBe, MaxInt, z)
	fmt.Println(ToBe, MaxInt, z)

	var hello int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", hello, f, b, s) // 0 0 false "" (zero values) for each type

	// Type conversions
	var x, y int = 100, 202
	var ff float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(ff)
	fmt.Println(x, y, z)

	// Type inference

	var tao int
	hid := tao // type of hid is inferred from tao
	fmt.Printf("hid is of type %T\n", hid)

	v := 42 // type of v is inferred from the right hand side
	fmt.Printf("v is of type %T\n", v)
	fl := 3.142 // type of fl is inferred from the right hand side
	fmt.Printf("fl is of type %T\n", fl)
	comp := 0 + 0i // type of comp is inferred from the right hand side
	fmt.Printf("comp is of type %T\n", comp)

	// Constants
	fmt.Println(Pi)

	// Numeric Constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
