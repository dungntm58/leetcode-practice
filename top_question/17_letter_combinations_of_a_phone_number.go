package top_question

/// Letter Combinations of a Phone Number
// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
// Return the answer in any order.
// A mapping of digit to letters (just like on the telephone buttons) is given below.
// Note that 1 does not map to any letters.
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
// Input: digits = ""
// Output: []
// Input: digits = "2"
// Output: ["a","b","c"]

func LetterCombinations(digits string) []string {
	chars := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}
	var res []string
	if len(digits) == 0 {
		return res
	}
	var tmp []byte
	dfsLetterCombinations(digits, 0, chars, &tmp, &res)
	return res
}

func dfsLetterCombinations(digits string, index int, chars [][]byte, tmp *[]byte, res *[]string) {
	if index == len(digits) {
		*res = append(*res, string(*tmp))
		return
	}
	for _, c := range chars[digits[index]-'2'] {
		*tmp = append(*tmp, c)
		dfsLetterCombinations(digits, index+1, chars, tmp, res)
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}
