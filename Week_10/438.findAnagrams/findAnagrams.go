package main

//滑动窗口
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	slen := len(s)
	plen := len(p)
	for i := 0; i < slen-plen+1; i++ {
		if isAnagram(s, p, i, i+plen-1) {
			res = append(res, i)
		}
	}
	return res
}

func isAnagram(s string, p string, i, j int) bool {
	c := [26]int{}
	for x := i; x <= j; x++ {
		sch := s[x] - 'a'
		pch := p[x-i] - 'a'
		c[sch]++
		c[pch]--
	}

	for _, ch := range c {
		if ch != 0 {
			return false
		}
	}
	return true

}
