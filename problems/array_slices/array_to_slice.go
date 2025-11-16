// Write a function that:
//
// Initializes a fixed-size array of integers with 5 elements.
//
// Converts it into a slice.
//
// Appends new elements to the slice.
//
// Prints both the original array and the modified slice.

package main

import (
	"fmt"
)

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Original Array: len=%d, cap=%d, data=%v\n", len(nums), cap(nums), nums)
	new_nums_slice := nums[:]
	fmt.Println("nums has been converted to slice")
	fmt.Println("appending new elements")
	new_nums_slice = append(new_nums_slice,6,7)
	fmt.Printf("Slice after append: len=%d, cap=%d, data=%v\n", len(new_nums_slice), cap(new_nums_slice), new_nums_slice)
}
