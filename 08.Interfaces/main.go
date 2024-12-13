// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Abser interface {
// 	Abs() float64
// }

// func main() {
// 	var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}

// 	a = f  // a MyFloat implements Abser
// 	fmt.Println(a.Abs())
// 	a = &v // a *Vertex implements Abser

// 	// In the following line, v is a Vertex (not *Vertex)
// 	// and does NOT implement Abser.
// 	//a = v

// 	fmt.Println(a.Abs())
// }

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

//** Interfaces are implemented implicitly **//
// package main

// import "fmt"

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// }

// func (t T) M() {
// 	fmt.Println(t.S)
// }

// func main() {
// 	var i I = T{"hello"}
// 	i.M()
// }

//** Interface values **//

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// }

// func (t *T) M() {
// 	fmt.Println(t.S)
// }

// type F float64

// func (f F) M() {
// 	fmt.Println(f)
// }

// func main() {
// 	var i I

// 	i = &T{"Hello"}
// 	describe(i)
// 	i.M()

// 	i = F(math.Pi)
// 	describe(i)
// 	i.M()
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

//** Interface values with nil underlying values **//
// package main

// import "fmt"

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// }

// func (t *T) M() {
// 	if t == nil {
// 		fmt.Println("<nil>")
// 		return
// 	}
// 	fmt.Println(t.S)
// }

// func main() {
// 	var i I

// 	var t *T
// 	i = t
// 	describe(i)
// 	i.M()

// 	i = &T{"hello"}
// 	describe(i)
// 	i.M()
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

//** Nil interface values **//

// package main

// import "fmt"

// type I interface {
// 	M()
// }

// func main() {
// 	var i I
// 	describe(i)
// 	i.M() // panic: runtime error: invalid memory address or nil pointer dereference
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

//** The empty interface **//

// package main

// import "fmt"

// func main() {
// 	var i interface{}
// 	describe(i)

// 	i = 42
// 	describe(i)

// 	i = "hello"
// 	describe(i)
// }

// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

//** Type assertions **//

// package main

// import "fmt"

// func main() {
// 	var i interface{} = "Taohid"
// 	t := i.(string)
// 	fmt.Println(t)

// 	t, ok := i.(string)
// 	fmt.Println(t, ok)

// 	f, ok := i.(float64)
// 	fmt.Println(f, ok)

// 	b,_ := i.(uint)
// 	fmt.Println(b)

// 	f = i.(float64) // panic: interface conversion: interface {} is string, not float64
// 	fmt.Println(f)
// }

//** Type switches **//

// package main

// import "fmt"

// func checkType(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println("Integer and value is ", v)
// 	case string:
// 		fmt.Println("String and value is ", v)
// 	case float64:
// 		fmt.Println("Float and value is ", v)
// 	default:
// 		fmt.Printf("Unknown type %T\n", v)

// 	}
// }

// func main() {
// 	checkType(21)
// 	checkType("Hello")
// 	checkType(3.14)
// 	checkType(true)
// }
