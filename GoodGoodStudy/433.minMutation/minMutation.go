package main

var mutaMap map[uint8][]string = map[uint8][]string{
	'A': []string{"C", "G", "T"},
	'C': []string{"A", "G", "T"},
	'G': []string{"A", "C", "T"},
	'T': []string{"A", "C", "G"},
}

func idxOf(curStr string, bank []string) int {
	for idx, tmpStr := range bank {
		if tmpStr == curStr {
			return idx
		}
	}

	return -1
}

//BFS 广度优先搜索
func minMutation(start string, end string, bank []string) int {

	isUsed := make([]bool, len(bank))
	count := 0
	//目标不在基因库
	if idxOf(end, bank) == -1 {
		return -1
	}
	queue := []string{start}
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			curr := queue[i]
			if curr == end {
				return count
			}
			for j := 0; j < len(curr); j++ {
				for _, s := range mutaMap[curr[j]] {
					if idx := idxOf(curr[:j]+s+curr[j+1:], bank); idx != -1 && !isUsed[idx] {
						queue = append(queue, bank[idx])
						isUsed[idx] = true
					}
				}

			}
		}

		count++
		queue = queue[l:]
	}
	return -1
}
