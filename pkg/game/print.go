package game

import (
	"fmt"
	"strings"
	"time"
)

const (
	// optimalLineLengh is the optimal length of a line in a paragraph.
	optimalLineLengh = 80
	// typeWriteTimeout determines the speed at which the typewriter types.
	typeWriteTimeout = 5 * time.Millisecond
)

// typeWrite prints to the terminal in a typewriter animation.
func typeWrite(str string) {
	column := 0
	for _, r := range str {
		time.Sleep(typeWriteTimeout)
		s := string(r)
		if column > optimalLineLengh && s == " " {
			fmt.Println(s)
			column = 0
			continue
		}
		fmt.Print(s)
		column++
	}

	fmt.Println("")
}

// withIndentation adds indentation to a string.
func withIndentation(str string) string {
	return strings.ReplaceAll(str, "\n", "\n                 ")
}
