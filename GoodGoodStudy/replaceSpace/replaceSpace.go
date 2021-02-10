package main

import "strings"

//1.思路一：直接遍历
func replaceSpace(s string) string {
	res := ""
	for _, ch := range s {
		if ch == ' ' {
			res += "%20"
		} else {
			res += string(ch)
		}
	}

	return res
}

//2.思路二：stings包
func replaceSpace_2(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

//3.思路三：
func replaceSpace_3(s string) string {
	var res strings.Builder
	for _, ch := range s {
		if ch == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteRune(ch)
		}
	}
	return res.String()
}
