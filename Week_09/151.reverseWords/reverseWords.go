package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	newStr := cleanSpaces(strs)

	left := 0
	right := len(newStr) - 1
	for left < right {
		newStr[left], newStr[right] = newStr[right], newStr[left]
		left++
		right--
	}
	return strings.Join(newStr, " ")
}

func cleanSpaces(s []string) []string {
	i := 0
	j := 0
	length := len(s)
	for j < length {
		for j < length && (s[j] == " " || s[j] == "") {
			j++
		}
		for j < length && s[j] != "" && s[j] != " " {
			s[i] = s[j]
			i++
			j++
		}
	}
	return s[0:i]
}

func reverseWords_2(s string) string {
	var res string
	length := len(s)
	left, right := length-1, length-1
	for left >= 0 {
		for left >= 0 && s[left] != byte(' ') {
			left--
		}
		res += s[left+1:right+1] + " "
		for left >= 0 && s[left] == byte(' ') {
			left--
		}
		right = left
	}
	return strings.TrimSpace(res)
}

func main() {
	str := "hello world "
	fmt.Println(reverseWords(str))
}
