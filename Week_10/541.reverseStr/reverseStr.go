package main

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		swap(bytes, i, i+k-1)
	}
	return string(bytes)
}

func swap(bytes []byte, i, j int) {
	if j >= len(bytes) {
		j = len(bytes) - 1
	}
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
}
