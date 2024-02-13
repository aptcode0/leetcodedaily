package feb

// Problem: https://leetcode.com/problems/palindromic-substrings/description
/*
Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
*/

func countSubstrings(s string) int {
	count := 0
	for i, _ := range s {
		count += pal(i, i, s)
		count += pal(i, i+1, s)
	}
	return count
}

func pal(i, j int, s string) int {
	count := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
		count++
	}
	return count
}