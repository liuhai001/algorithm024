package main

func reverseStr(s string, k int) string {
	i := 0
	schar := []byte(s)
	n := len(schar)
	for i < n {
		j := min(i+k-1, n-1)
		swap(schar, i, j)
		i += 2 * k
	}

	return string(schar)
}

func swap(schar []byte, i, j int) {
	for i < j {
		tmp := schar[i]
		schar[i] = schar[j]
		schar[j] = tmp
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
