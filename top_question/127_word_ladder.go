package top_question

/// Word Ladder
// A transformation sequence from word beginWord to word endWord
// using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:
// - Every adjacent pair of words differs by a single letter.
// - Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
// - sk == endWord
// Given two words, beginWord and endWord, and a dictionary wordList,
// return the number of words in the shortest transformation sequence from beginWord to endWord,
// or 0 if no such sequence exists.
// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
// Output: 5
// Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.
// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
// Output: 0
// Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.

func LadderLength(beginWord string, endWord string, wordList []string) int {
	allComboMap := map[string][]string{}
	for _, w := range wordList {
		for i := 0; i < len(w); i++ {
			newWord := w[:i] + "*" + w[i+1:]
			allComboMap[newWord] = append(allComboMap[newWord], w)
		}
	}
	type pair struct {
		value string
		level int
	}
	queue := []pair{{beginWord, 1}}
	visited := map[string]bool{beginWord: true}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 0; i < len(cur.value); i++ {
			newWord := cur.value[:i] + "*" + cur.value[i+1:]
			for _, adjacentWord := range allComboMap[newWord] {
				if adjacentWord == endWord {
					return cur.level + 1
				}
				if _, ok := visited[adjacentWord]; !ok {
					visited[adjacentWord] = true
					queue = append(queue, pair{adjacentWord, cur.level + 1})
				}
			}
		}
	}
	return 0
}
