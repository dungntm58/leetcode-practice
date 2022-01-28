package top_question

/// Design Add and Search Words Data Structure
// Design a data structure that supports adding new words and finding if a string matches any previously added string.
// Implement the WordDictionary class:
// - WordDictionary() Initializes the object.
// - void addWord(word) Adds word to the data structure, it can be matched later.
// - bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise.
// 	word may contain dots '.' where dots can be matched with any letter.

type WordDictionary struct {
	root *TrieNode
}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{root: NewTrieNode(false)}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for i, c := range word {
		if curr.children[c-'a'] == nil {
			curr.children[c-'a'] = NewTrieNode(i == len(word)-1)
		}
		curr = curr.children[c-'a']
	}
}

func (this *WordDictionary) Search(word string) bool {
	return searchOnTrieNodeWithTheRestLetters([]rune(word), this.root)
}

func searchOnTrieNodeWithTheRestLetters(runes []rune, node *TrieNode) bool {
	if len(runes) == 0 {
		return node.isEndOfWord
	}
	if runes[0] == '.' {
		for _, child := range node.children {
			if child != nil && searchOnTrieNodeWithTheRestLetters(runes[1:], child) {
				return true
			}
		}
	} else if node.children[runes[0]-'a'] != nil {
		return searchOnTrieNodeWithTheRestLetters(runes[1:], node.children[runes[0]-'a'])
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
