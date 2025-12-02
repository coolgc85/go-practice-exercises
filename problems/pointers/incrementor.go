package main

import (
	"fmt"
)

// Increment adds one to the integer value pointed to by n.
// It demonstrates how pointers can be used to modify a variable's
// value from within a function.
func Increment(n *int) {
	*n++ // Idiomatic way to increment the value at the pointer's address
}

// incrementByValue attempts to increment an integer, but it only
// modifies a local copy of the variable. This demonstrates Go's
// pass-by-value nature.
func incrementByValue(n int) {
	n++
	// This function intentionally does not print. Functions should generally
	// avoid side effects like printing unless it's their primary purpose.
}

func main() {
	// --- Demonstration of Increment (pass-by-pointer) ---
	number := 6
	fmt.Println("--- Demonstrating Pass-by-Pointer ---")
	fmt.Printf("Original number: %d\n", number)

	Increment(&number)
	fmt.Printf("After first Increment: %d\n", number)

	Increment(&number)
	fmt.Printf("After second Increment: %d\n", number)
	fmt.Println() // Add a newline for better formatting

	// --- Demonstration of incrementByValue (pass-by-value) ---
	fmt.Println("--- Demonstrating Pass-by-Value ---")
	fmt.Printf("Number before incrementByValue call: %d\n", number)
	incrementByValue(number)
	fmt.Printf("Number after incrementByValue call: %d (Note: it's unchanged)\n", number)
}