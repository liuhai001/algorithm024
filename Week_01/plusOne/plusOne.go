package main

import "fmt"

func plusOne(digits []int) []int {
	//1、把digits 转换成数字， bug，溢出怎么办？
	//2、实现一个加法规则

	length := len(digits)
	var i int
	for i = length - 1; i >= 0; i-- {
		if digits[i]+1 > 9 {
			digits[i] = 0
			continue
		}

		digits[i] += 1
		return digits
	}

	new_digits := make([]int, 0, length+1)
	if i == -1 {
		new_digits = append([]int{1}, digits...)
	}

	return new_digits
}

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))
}
