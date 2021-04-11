package main

//空间复杂度O(1) 时间复杂度O(n^2)
func longestPalindrome(s string) string {
	res := ""
	length := len(s)
	for i := 0; i < length; i++ {
		//以s[i]为中心的回文子串
		s1 := palindrome(s, i, i)
		//以s[i]和s[i+1]为中心的回文子串
		s2 := palindrome(s, i, i+1)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}

	}
	return res
}

func palindrome(s string, left, right int) string {
	//防止索引越界
	for left >= 0 && right < len(s) && s[left] == s[right] {
		//中心向两边展开
		left--
		right++
	}

	return s[left+1 : right]
}
