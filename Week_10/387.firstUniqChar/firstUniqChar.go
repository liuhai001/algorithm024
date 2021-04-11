package main

//两次遍历, map, 时间复杂度O(n) 空间复杂度O(1)
func firstUniqChar(s string) int {
	cnt := [26]int{}
	for _,c := range s{
		cnt[c-'a']++
	}

	for i,c := range s{
		if cnt[c-'a'] == 1{
			return i
		}
	}
	return -1
}
