package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		println("no argument")
		return 0
	}

	var words []string = strings.Fields(s)
	println(words[len(words)-1])
	return len(words[len(words)-1])
}

func main() {
	var result int = lengthOfLastWord("Not my biz           ")
	fmt.Println(result)

	result = lengthOfLastWord("  fdsfs   ")
	fmt.Println(result)

}
