package main

import (
	"fmt"
	"go-practice-exercises/problems/package/mathutils"
)

func main() {

	sumResult := mathutils.Add(2, 3)
	fmt.Printf("Sum Result: %d\n", sumResult)

	diffResult := mathutils.Subtract(10, 8)
	fmt.Printf("Difference Result: %d\n", diffResult)

}