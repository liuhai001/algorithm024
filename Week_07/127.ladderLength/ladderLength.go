package main

//bfs 层序遍历
func ladderLength_1(beginWord string, endWord string, wordList []string) int {
	isInList := make(map[string]bool)
	count := 1
	//初始化
	for _, word := range wordList {
		isInList[word] = true
	}
	//目标不在基因库
	if !isInList[endWord] {
		return 0
	}
	visited := make(map[string]bool, len(wordList))
	visited[beginWord] = true

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
					if isInList[tmp]&& !visited[tmp]{
						queue = append(queue, tmp)
						visited[tmp] = true
					}
				}

			}
		}

		count++
		queue = queue[length:]
	}
	return 0
}

//双向bfs
func ladderLength_2(beginWord string, endWord string, wordList []string) int {
	if wordList == nil || len(wordList) == 0 {
		return 0
	}
	if endWord == "" || beginWord == "" {
		return 0
	}

	myMap := make(map[string]bool, len(wordList))
	for _, str := range wordList {
		myMap[str] = true
	}
	if _, ok := myMap[endWord]; !ok {
		return 0
	}
	delete(myMap, beginWord)

	visited := make(map[string]bool, len(wordList))
	visited[beginWord] = true

	//从起点和终点双向搜索
	beginVisited, endVisited := map[string]bool{beginWord: true}, map[string]bool{endWord: true}
	step := 1

	for len(beginVisited) > 0 && len(endVisited) > 0 {
		//从小的那一方向搜
		if len(beginVisited) > len(endVisited) {
			beginVisited, endVisited = endVisited, beginVisited
		}

		newVisited := map[string]bool{}
		for word := range beginVisited {
			for i := 0; i < len(word); i++ {
				for ch := 'a'; ch <= 'z'; ch++ {
					if rune(word[i]) == ch {
						continue
					}
					newWord := word[:i] + string(ch) + word[i+1:]
					if _, ok := myMap[newWord]; ok {
						if _, ok := endVisited[newWord]; ok {
							//找到了
							step++
							return step
						}

						if _, yes := visited[newWord]; !yes {
							visited[newWord] = true
							newVisited[newWord] = true
						}
					}
				}
			}
		}
		beginVisited = newVisited
		step++
	}

	return 0
}
