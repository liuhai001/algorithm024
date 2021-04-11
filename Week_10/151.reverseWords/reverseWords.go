package main

import "strings"

func reverseWords(s string) string {
	list := strings.Split(s, " ")
	arr := make([]string, 0)
	for i := len(list) - 1; i >= 0; i-- {
		if len(list[i]) > 0 {
			arr = append(arr, list[i])
		}
	}
	return strings.Join(arr, " ")
}
