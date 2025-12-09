package main

import "fmt"

const rangeConstant = 10

// useResource method simulates a resource usage by given ID
func useResource(i int) {
	fmt.Printf("opening resource %d\n", i)

	//simulate computation
	for range rangeConstant {
	}
}

//releaseResource method simulates the resource release
func releaseResource(i int) {
	fmt.Printf("closing resource %d\n", i)
}

//main method is the entry point
func main() {

	for i := range rangeConstant {
		useResource(i)
		defer releaseResource(i)
	}

}
