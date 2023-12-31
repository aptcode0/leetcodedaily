package dec

// Problem: https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/description
/*
Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
Output: true
Explanation:
word1 represents string "ab" + "c" -> "abc"
word2 represents string "a" + "bc" -> "abc"
The strings are the same, so return true.
*/

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	l1, l2 := 0, 0
	i1, i2 := 0, 0
	for l1 < len(word1) && l2 < len(word2) {
		if i1 >= len(word1[l1]) {
			l1++
			i1 = 0
		}
		if i2 >= len(word2[l2]) {
			l2++
			i2 = 0
		}
		if l1 >= len(word1) || l2 >= len(word2) {
			break
		}

		if word1[l1][i1] != word2[l2][i2] {
			return false
		}
		i1++
		i2++
	}

	return l1 == len(word1) && l2 == len(word2)
}
