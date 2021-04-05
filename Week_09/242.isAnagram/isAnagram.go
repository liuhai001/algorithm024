package main

import "sort"

//排序
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s1 := []byte(s)
	s2 := []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

//hash表
func isAnagram_1(s string, t string) bool {
	c1 := [26]int{}
	c2 := [26]int{}
	for _, c := range s {
		c1[c-'a']++
	}
	for _, c := range t {
		c2[c-'a']++
	}
	return c1 == c2
}
