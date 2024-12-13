package main

import "fmt"

// Stringer is an interface that has a single method called String which returns a string.
// fmt package calls this method to get the default string representation of the type.
// A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

// type Stringer interface {
// 	String() string
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }

// func main() {
// 	a := Person{"Avik", 25}
// 	b := Person{"Sizan", 27}
// 	fmt.Println(a, b)
// }


//** Exercise: Stringers **//
// Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v",i[0],i[1],i[2],i[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}