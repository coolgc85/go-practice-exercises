package main

import "fmt"

func gradeIt(score float64) string {
	var message string
	if score < 7 {
		message = fmt.Sprintf("Fail with %.2f", score)
	} else {
		switch {
		case score >= 9:
			message = fmt.Sprintf("Excellent %.2f", score)
		case score >= 7:
			message = fmt.Sprintf("Pass with %.2f", score)
		}
	}
	return message
}

func main() {
	score1 := 10.0
	score2 := 7.5
	score3 := 6.7

	fmt.Println("Test with an Excellent score")
	message := gradeIt(score1)
	fmt.Println(message)
	fmt.Println("==========================")

	fmt.Println("Test with Pass")
	message = gradeIt(score2)
	fmt.Println(message)
	fmt.Println("==========================")


	fmt.Println("Test with Fail")
	message = gradeIt(score3)
	fmt.Println(message)
	fmt.Println("==========================")


}
