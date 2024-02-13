package feb

// Problem: https://leetcode.com/problems/find-first-palindromic-string-in-the-array/description
/*
Input: words = ["abc","car","ada","racecar","cool"]
Output: "ada"
Explanation: The first string that is palindromic is "ada".
Note that "racecar" is also palindromic, but it is not the first.
*/

func firstPalindrome(words []string) string {
	for _, s := range words {
		if isPali(s) {
			return s
		}
	}
	return ""
}

func isPali(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
