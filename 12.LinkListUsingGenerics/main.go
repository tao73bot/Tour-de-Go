package main

import (
	"fmt"
)

// List represents a singly-linked list with generic type T.
type List[T comparable] struct {
	value T
	next  *List[T]
}

// NewList creates a new singly-linked list with the provided values.
func NewList[T comparable](values ...T) *List[T] {
	if len(values) == 0 {
		return nil
	}
	head := &List[T]{value: values[0]}
	current := head
	for _, v := range values[1:] {
		current.next = &List[T]{value: v}
		current = current.next
	}
	return head
}

// Append adds a new value to the end of the list.
func (l *List[T]) Append(value T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{value: value}
}

// Print prints all the values in the list.
func (l *List[T]) Print() {
	current := l
	for current != nil {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

// Length returns the number of elements in the list.
func (l *List[T]) Length() int {
	count := 0
	current := l
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// Find returns true if the list contains the specified value.
func (l *List[T]) Find(value T) bool {
	current := l
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func main() {
	// Create a new list with integers
	list := NewList(1, 2, 3, 4, 5)
	list.Print()

	// Append a new value to the list
	list.Append(6)
	list.Print()

	// Print the length of the list
	fmt.Println("Length:", list.Length())

	// Find a value in the list
	fmt.Println("Find 3:", list.Find(3))
	fmt.Println("Find 7:", list.Find(7))
}
