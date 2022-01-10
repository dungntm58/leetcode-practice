package contest

import "strings"

// 2114
/// Maximum Number of Words Found in Sentences
// A sentence is a list of words that are separated by a single space with no leading or trailing spaces.
// You are given an array of strings sentences, where each sentences[i] represents a single sentence.
// Return the maximum number of words that appear in a single sentence.
// Input: sentences = ["alice and bob love leetcode", "i think so too", "this is great thanks very much"]
// Output: 6
// Explanation:
// - The first sentence, "alice and bob love leetcode", has 5 words in total.
// - The second sentence, "i think so too", has 4 words in total.
// - The third sentence, "this is great thanks very much", has 6 words in total.
// Thus, the maximum number of words in a single sentence comes from the third sentence, which has 6 words.
// Input: sentences = ["please wait", "continue to fight", "continue to win"]
// Output: 3
// Explanation: It is possible that multiple sentences contain the same number of words.
// In this example, the second and third sentences (underlined) have the same number of words.

func MostWordsFound(sentences []string) int {
	max := 0
	for _, s := range sentences {
		wordCount := len(strings.Fields(s))
		if wordCount > max {
			max = wordCount
		}
	}
	return max
}

// func MostWordsFound(sentences []string) int {
// 	max := 0
// 	for _, s := range sentences {
// 		spaceCount := 1
// 		for _, c := range s {
// 			if c == ' ' {
// 				spaceCount++
// 			}
// 		}
// 		if spaceCount > max {
// 			max = spaceCount
// 		}
// 	}
// 	return max
// }

// 2115
/// Find All Possible Recipes from Given Supplies
// You have information about n different recipes. You are given a string array recipes and a 2D string array ingredients.
// The ith recipe has the name recipes[i], and you can create it if you have all the needed ingredients from ingredients[i].
// Ingredients to a recipe may need to be created from other recipes,
// i.e., ingredients[i] may contain a string that is in recipes.
// You are also given a string array supplies containing all the ingredients that you initially have,
// and you have an infinite supply of all of them.
// Return a list of all the recipes that you can create. You may return the answer in any order.
// Note that two recipes may contain each other in their ingredients.
// Input: recipes = ["bread"], ingredients = [["yeast","flour"]], supplies = ["yeast","flour","corn"]
// Output: ["bread"]
// Explanation:
// We can create "bread" since we have the ingredients "yeast" and "flour".
// Input: recipes = ["bread","sandwich"], ingredients = [["yeast","flour"],["bread","meat"]], supplies = ["yeast","flour","meat"]
// Output: ["bread","sandwich"]
// Explanation:
// We can create "bread" since we have the ingredients "yeast" and "flour".
// We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
// Input: recipes = ["bread","sandwich","burger"], ingredients = [["yeast","flour"],["bread","meat"],["sandwich","meat","bread"]], supplies = ["yeast","flour","meat"]
// Output: ["bread","sandwich","burger"]
// Explanation:
// We can create "bread" since we have the ingredients "yeast" and "flour".
// We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
// We can create "burger" since we have the ingredient "meat" and can create the ingredients "bread" and "sandwich".

func FindAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	return nil
}

// 2116
/// Check if a Parentheses String Can Be Valid
// A parentheses string is a non-empty string consisting only of '(' and ')'.
// It is valid if any of the following conditions is true:
// It is ().
// It can be written as AB (A concatenated with B), where A and B are valid parentheses strings.
// It can be written as (A), where A is a valid parentheses string.
// You are given a parentheses string s and a string locked,
// both of length n. locked is a binary string consisting only of '0's and '1's. For each index i of locked,
// If locked[i] is '1', you cannot change s[i].
// But if locked[i] is '0', you can change s[i] to either '(' or ')'.
// Return true if you can make s a valid parentheses string. Otherwise, return false.
// Input: s = "))()))", locked = "010100"
// Output: true
// Explanation: locked[1] == '1' and locked[3] == '1', so we cannot change s[1] or s[3].
// We change s[0] and s[4] to '(' while leaving s[2] and s[5] unchanged to make s valid.
// Input: s = "()()", locked = "0000"
// Output: true
// Explanation: We do not need to make any changes because s is already valid.
// Input: s = ")", locked = "0"
// Output: false
// Explanation: locked permits us to change s[0].
// Changing s[0] to either '(' or ')' will not make s valid.

func CanBeValid(s string, locked string) bool {
	return false
}

// 2117
/// Abbreviating the Product of a Range
// You are given two positive integers left and right with left <= right.
// Calculate the product of all integers in the inclusive range [left, right].
// Since the product may be very large, you will abbreviate it following these steps:
// 1. Count all trailing zeros in the product and remove them. Let us denote this count as C.
// - For example, there are 3 trailing zeros in 1000, and there are 0 trailing zeros in 546.
// 2. Denote the remaining number of digits in the product as d.
// If d > 10, then express the product as <pre>...<suf> where <pre> denotes the first 5 digits of the product,
// and <suf> denotes the last 5 digits of the product after removing all trailing zeros. If d <= 10, we keep it unchanged.
// - For example, we express 1234567654321 as 12345...54321, but 1234567 is represented as 1234567.
// 3. Finally, represent the product as a string "<pre>...<suf>eC".
// - For example, 12345678987600000 will be represented as "12345...89876e5".
// Return a string denoting the abbreviated product of all integers in the inclusive range [left, right].
// Input: left = 1, right = 4
// Output: "24e0"
// Explanation: The product is 1 × 2 × 3 × 4 = 24.
// There are no trailing zeros, so 24 remains the same. The abbreviation will end with "e0".
// Since the number of digits is 2, which is less than 10, we do not have to abbreviate it further.
// Thus, the final representation is "24e0".
// Input: left = 2, right = 11
// Output: "399168e2"
// Explanation: The product is 39916800.
// There are 2 trailing zeros, which we remove to get 399168. The abbreviation will end with "e2".
// The number of digits after removing the trailing zeros is 6, so we do not abbreviate it further.
// Hence, the abbreviated product is "399168e2".
// Input: left = 371, right = 375
// Output: "7219856259e3"
// Explanation: The product is 7219856259000.

func AbbreviateProduct(left int, right int) string {
	return ""
}
