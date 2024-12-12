package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// A map maps keys to values.
	// initialize a map with the built-in make function
	m := make(map[string]int)
	m["Avik"] = 1
	fmt.Println("The value:", m["Avik"])

	m["Fahad"] = 2
	fmt.Println("The value:", m["Fahad"])

	x := m["Avik"]
	fmt.Println("The value:", x)

	// delete a key
	delete(m, "Avik")
	fmt.Println("The value:", m["Avik"])

	v, ok := m["Avik"]
	fmt.Println("The value:", v, "Present?", ok)
}