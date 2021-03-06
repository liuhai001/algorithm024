package main

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	isInList := make(map[string]bool)
	res := make([][]string,0)
	//初始化
	for _, word := range wordList {
		isInList[word] = true
	}
	//目标不在基因库
	if !isInList[endWord] {
		return res
	}

	queue := []string{beginWord}
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			curr := queue[i]
			if curr == endWord {
				return count
			}
			for j := 0; j < len(curr); j++ {
				for s := 'a'; s <= 'z'; s++ {
					tmp := curr[:j] + string(s) + curr[j+1:]
					if isInList[tmp] {
						queue = append(queue, tmp)
						isInList[tmp] = false
					}
				}

			}
		}

		count++
		queue = queue[length:]
	}
	return res
}