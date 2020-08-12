package game

import (
	"fmt"
	"strings"
	"time"
)

// optimalLineLengh is the optimal length of a line in a paragraph.
const (
	optimalLineLengh = 80
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

// printIndent prints an indented text.
func printIndent(str string) {
	str = strings.ReplaceAll(str, "\n", "\n                 ")
	fmt.Print(str)
}
