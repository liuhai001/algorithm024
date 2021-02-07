package main

import (
	"fmt"
	"sort"
)

//时间复杂度 O(nklogk) 空间复杂度 n
func groupAnagrams_1(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, str := range strs {
		s1 := []byte(str)
		sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
		mp[string(s1)] = append(mp[string(s1)], str)
	}
	res := make([][]string, 0)
	for _, strs := range mp {
		res = append(res, strs)
	}
	return res
}

//时间复杂度 O(nk) 空间复杂度 nk
func groupAnagrams_2(strs []string) [][]string {
	mp := make(map[[26]int][]string)
	for _, str := range strs {
		var key [26]int
		for _, ch := range str {
			key[ch-'a']++
		}
		mp[key] = append(mp[key], str)
	}
	res := make([][]string, 0)
	for _, strs := range mp {
		res = append(res, strs)
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams_2(strs))
}
