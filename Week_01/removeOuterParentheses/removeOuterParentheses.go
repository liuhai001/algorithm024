package main

import (
	"fmt"
	"strings"
)

func removeOuterParentheses(S string) string {
	var sb strings.Builder
	level := 0
	for _, ch := range S {
		if ch == ')' {
			level--
		}
		if level >= 1 {
			sb.WriteByte(byte(ch))
		}
		if ch == '(' {
			level++
		}
	}

	return sb.String()
}

func main()  {
	str := "(()())(())"
	fmt.Println(removeOuterParentheses(str))
}