package top_question

/// Stream of Characters
// Design an algorithm that accepts a stream of characters
// and checks if a suffix of these characters is a string of a given array of strings words.
// For example, if words = ["abc", "xyz"] and the stream added the four characters (one by one) 'a', 'x', 'y', and 'z',
// your algorithm should detect that the suffix "xyz" of the characters "axyz" matches "xyz" from words.
// Implement the StreamChecker class:
// - StreamChecker(String[] words) Initializes the object with the strings array words.
// - boolean query(char letter) Accepts a new character from the stream
// and returns true if any non-empty suffix from the stream forms a word that is in words.
// Input
// ["StreamChecker", "query", "query", "query", "query", "query", "query", "query", "query", "query", "query", "query", "query"]
// [[["cd", "f", "kl"]], ["a"], ["b"], ["c"], ["d"], ["e"], ["f"], ["g"], ["h"], ["i"], ["j"], ["k"], ["l"]]
// Output
// [null, false, false, false, true, false, true, false, false, false, false, false, true]
// Explanation
// StreamChecker streamChecker = new StreamChecker(["cd", "f", "kl"]);
// streamChecker.query("a"); // return False
// streamChecker.query("b"); // return False
// streamChecker.query("c"); // return False
// streamChecker.query("d"); // return True, because 'cd' is in the wordlist
// streamChecker.query("e"); // return False
// streamChecker.query("f"); // return True, because 'f' is in the wordlist
// streamChecker.query("g"); // return False
// streamChecker.query("h"); // return False
// streamChecker.query("i"); // return False
// streamChecker.query("j"); // return False
// streamChecker.query("k"); // return False
// streamChecker.query("l"); // return True, because 'kl' is in the wordlist

type StreamChecker struct {
	trie      *Trie
	queries   []rune
	maxLength int
}

func ConstructorStreamChecker(words []string) StreamChecker {
	trie, maxLength := NewTrie(), 0
	for _, w := range words {
		runes := []rune(w)
		reverseRunes(runes)
		trie.Insert(string(runes))
		if len(runes) > maxLength {
			maxLength = len(runes)
		}
	}
	return StreamChecker{trie: trie, maxLength: maxLength}
}

func (this *StreamChecker) Query(letter byte) bool {
	if len(this.queries) < this.maxLength {
		this.queries = append(this.queries, '0')
	}
	for i := len(this.queries) - 1; i > 0; i-- {
		this.queries[i] = this.queries[i-1]
	}
	this.queries[0] = rune(letter)

	curr := this.trie.root
	for _, c := range this.queries {
		if curr == nil {
			return false
		}
		if curr.isEndOfWord {
			return true
		}
		curr = curr.children[c-'a']
	}
	if curr == nil {
		return false
	}
	return curr.isEndOfWord
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
