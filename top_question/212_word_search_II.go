package top_question

/// Given an m x n board of characters and a list of strings words, return all words on the board.
// Each word must be constructed from letters of sequentially adjacent cells,
// where adjacent cells are horizontally or vertically neighboring.
// The same letter cell may not be used more than once in a word.
// Input:
// board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]],
// words = ["oath","pea","eat","rain"]
// Output: ["eat","oath"]
// Input:
// board = [["a","b"],["c","d"]],
// words = ["abcb"]
// Output: []

func FindWords(board [][]byte, words []string) []string {
	trie := NewTrie()
	for _, w := range words {
		trie.Insert(w)
	}
	m := len(board)
	n := len(board[0])
	r := []string{}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			findWordsDfs(board, i, j, []rune{}, trie.GetRoot(), visited, &r)
		}
	}
	return r
}

func findWordsDfs(board [][]byte, i, j int, runeArr []rune, node *TrieNode, visited [][]bool, ans *[]string) {
	if visited[i][j] {
		return
	}
	char := board[i][j]
	node = node.Child(rune(char))
	if node == nil {
		return
	}
	visited[i][j] = true
	runeArr = append(runeArr, rune(char))
	if node.IsEndOfWord() {
		*ans = append(*ans, string(runeArr))
		node.MarkEndOfWordAsFalse()
	}
	if j > 0 {
		findWordsDfs(board, i, j-1, runeArr, node, visited, ans)
	}
	if i < len(board)-1 {
		findWordsDfs(board, i+1, j, runeArr, node, visited, ans)
	}
	if j < len(board[0])-1 {
		findWordsDfs(board, i, j+1, runeArr, node, visited, ans)
	}
	if i > 0 {
		findWordsDfs(board, i-1, j, runeArr, node, visited, ans)
	}
	visited[i][j] = false
}
