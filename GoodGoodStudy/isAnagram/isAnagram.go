package main

import (
	"fmt"
	"sort"
)

//1. hash 时间复杂度O(n) n是字符串长度
func isAnagram_1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var ch [26]int
	for i, _ := range s {
		ch[s[i]-'a']++
		ch[t[i]-'a']--
	}

	for _, val := range ch {
		if val != 0 {
			return false
		}
	}

	return true
}

//2.sort 时间复杂度O(nlogn) n是字符串长度
func isAnagram_2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s1 := []byte(s)
	s2 := []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram_1(s, t))
	fmt.Println(isAnagram_2(s, t))
}
