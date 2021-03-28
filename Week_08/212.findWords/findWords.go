package main

type TrieNode struct {
	Next []*TrieNode
	Word string
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Next: make([]*TrieNode, 26),
		Word: "",
	}

}

func buildTrie(words []string) *TrieNode {
	root := NewTrieNode()
	for _, w := range words {
		p := root
		for _, c := range w {
			pos := c - 'a'
			if p.Next[pos] == nil {
				p.Next[pos] = NewTrieNode()
			}
			p = p.Next[pos]
		}
		p.Word = w
	}
	return root
}

func findWords(board [][]byte, words []string) []string {
	res := make([]string, 0, len(words))
	//异常条件判断
	if len(words) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return res
	}
	root := buildTrie(words)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, root, &res)
		}
	}

	return res
}

func dfs(board [][]byte, i, j int, root *TrieNode, res *[]string) {
	c := board[i][j]
	//已遍历过，或者下一个字母不存在trie树中
	if c == '#' || root.Next[c-'a'] == nil {
		return
	}
	root = root.Next[c-'a']
	if root.Word != "" {
		//找到一个结果
		*res = append(*res, root.Word)
		root.Word = "" //去重
		//return //不能返回，还可能要继续寻找，比如 "oa"-> "oaa"
	}

	//继续寻找,标记已访问
	board[i][j] = '#'
	if i > 0 {
		dfs(board, i-1, j, root, res)
	}
	if j > 0 {
		dfs(board, i, j-1, root, res)
	}
	if i < len(board)-1 {
		dfs(board, i+1, j, root, res)
	}
	if j < len(board[0])-1 {
		dfs(board, i, j+1, root, res)
	}

	//恢复状态
	board[i][j] = c

}

func main() {
	buildTrie([]string{"oa", "oaa"})
}
