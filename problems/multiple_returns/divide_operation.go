package main

import (
	"fmt"
)

func divide(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}

	return numerator / denominator, nil
}

func main() {

	numerator := 5.0
	denominator := 8.0

	// --- Case 1: Successful Operation ---
	fmt.Printf("Dividing %.2f / %.2f\n", numerator, denominator)
	result, err := divide(numerator, denominator)

	if err != nil {
		fmt.Printf("Error %v\n", err)
	} else {
		fmt.Println("Divide result, denominator greater than 0 => ", result)
	}

	// --- Case 2: Operation that Fails ---
	fmt.Println("Updating denominator to => 0")
	denominator = 0
	fmt.Printf("Dividing %.2f / %.2f\n", numerator, denominator)
	result, err = divide(numerator, denominator)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	} else {
		fmt.Println("Divide result", result)
	}

}
