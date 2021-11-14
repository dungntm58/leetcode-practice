package top_question

import (
	"strconv"
)

/// Iterator for Combination
// Design the CombinationIterator class:
// - CombinationIterator(string characters, int combinationLength)
// Initializes the object with a string characters of sorted distinct lowercase English letters
// and a number combinationLength as arguments.
// - next() Returns the next combination of length combinationLength in lexicographical order.
// - hasNext() Returns true if and only if there exists a next combination.
// Input
// ["CombinationIterator", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
// [["abc", 2], [], [], [], [], [], []]
// Output
// [null, "ab", true, "ac", true, "bc", false]

// Explanation
// CombinationIterator itr = new CombinationIterator("abc", 2);
// itr.next();    // return "ab"
// itr.hasNext(); // return True
// itr.next();    // return "ac"
// itr.hasNext(); // return True
// itr.next();    // return "bc"
// itr.hasNext(); // return False

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

type CombinationIterator struct {
	characters           []rune
	combinationLength    int
	currentDecimal       uint64
	allCombinationLength uint64
	currentIndex         uint64
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	n := len(characters)
	return CombinationIterator{
		characters:           []rune(characters),
		combinationLength:    combinationLength,
		currentDecimal:       0,
		allCombinationLength: binomal(n, combinationLength),
		currentIndex:         0,
	}
}

func (this *CombinationIterator) Next() string {
	n := len(this.characters)

	var binary []rune = make([]rune, n)
	for i := 0; i < n; i++ {
		binary[i] = '0'
	}
	if this.currentIndex == 0 {
		for i := 0; i < this.combinationLength; i++ {
			binary[i] = '1'
		}
		this.currentDecimal, _ = strconv.ParseUint(string(binary), 2, 64)
	} else {
		for i := this.currentDecimal - 1; i > 0; i-- {
			r := []rune(strconv.FormatUint(i, 2))
			lenR := len(r)
			if containNBit(r, this.combinationLength) {
				for i, el := range r {
					binary[i+n-lenR] = el
				}
				this.currentDecimal = i
				break
			}
		}
	}

	r := make([]rune, this.combinationLength)
	j := 0
	for i := 0; i < n; i++ {
		if binary[i] == '1' {
			r[j] = this.characters[i]
			j++
		}
	}
	this.currentIndex++
	return string(r)
}

func (this *CombinationIterator) HasNext() bool {
	return this.currentIndex < this.allCombinationLength
}

func binomal(n int, k int) uint64 {
	if k > n/2 {
		k = n - k
	}
	b := uint64(1)
	for i := 1; i <= k; i++ {
		b = (uint64(n - k + i)) * b / uint64(i)
	}
	return b
}

func containNBit(arr []rune, k int) bool {
	r := 0
	for _, el := range arr {
		if el == '1' {
			r++
		}
	}
	return r == k
}
