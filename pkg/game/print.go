package game

import (
	"fmt"
	"strings"
	"time"
)

// optimalLineLengh is the optimal length of a line in a paragraph.
const optimalLineLengh = 80

// typeWrite prints to the terminal in a typewriter animation.
func typeWrite(str string) {
	column := 0
	for _, r := range str {
		time.Sleep(5 * time.Millisecond)
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
	str = strings.Replace(str, "\n", "\n                 ", -1)
	fmt.Print(str)
}
