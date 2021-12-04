package top_question

/// Implement Trie (Prefix Tree)
// A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store
// and retrieve keys in a dataset of strings.
// There are various applications of this data structure, such as autocomplete and spellchecker.
// Implement the Trie class:
// - Trie() Initializes the trie object.
// - void insert(String word) Inserts the string word into the trie.
// - boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before),
// and false otherwise.
// - boolean startsWith(String prefix)
// Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
// Input
// ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
// [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
// Output
// [null, null, true, false, true, null, true]
// Explanation
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // return True
// trie.search("app");     // return False
// trie.startsWith("app"); // return True
// trie.insert("app");
// trie.search("app");     // return True

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make([]*TrieNode, 26)}}
}

func ConstructorTrie() Trie {
	return Trie{root: &TrieNode{children: make([]*TrieNode, 26)}}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	for _, c := range word {
		if curr.children[c-'a'] == nil {
			curr.children[c-'a'] = NewTrieNode(c, false)
		}
		curr = curr.children[c-'a']
	}
	curr.isEndOfWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for _, c := range word {
		curr = curr.children[c-'a']
		if curr == nil {
			return false
		}
	}
	return curr.isEndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for _, c := range prefix {
		curr = curr.children[c-'a']
		if curr == nil {
			return false
		}
	}
	return true
}

type TrieNode struct {
	val         rune
	children    []*TrieNode
	isEndOfWord bool
}

func NewTrieNode(value rune, isEndOfWord bool) *TrieNode {
	return &TrieNode{val: value, children: make([]*TrieNode, 26), isEndOfWord: isEndOfWord}
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
