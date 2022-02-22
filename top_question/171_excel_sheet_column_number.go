package top_question

/// Excel Sheet Column Number
// Given a string columnTitle that represents the column title as appear in an Excel sheet,
// return its corresponding column number.
// For example:
// A -> 1
// B -> 2
// C -> 3
// ...
// Z -> 26
// AA -> 27
// AB -> 28
// ...
// Input: columnTitle = "A"
// Output: 1
// Input: columnTitle = "AB"
// Output: 28
// Input: columnTitle = "ZY"
// Output: 701

func TitleToNumber(columnTitle string) int {
	res := 0
	for i := 0; i < len(columnTitle); i++ {
		res = res*26 + int(columnTitle[i]-'A') + 1
	}
	return res
}
