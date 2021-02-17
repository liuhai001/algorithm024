package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	digitMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	recursion("", digits, 0, &res, digitMap)
	return res
}

func recursion(s string, digits string, level int, res *[]string, digitMap map[byte]string) {
	if level == len(digits) {
		*res = append(*res, s)
		return
	}

	for _, ch := range digitMap[digits[level]] {
		//进入下一层
		recursion(s+string(ch), digits, level+1, res, digitMap)
	}
}
