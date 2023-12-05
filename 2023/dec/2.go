package dec

// Problem: https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/description
/*
Input: words = ["cat","bt","hat","tree"], chars = "atach"
Output: 6
Explanation: The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.
*/

func countCharacters(words []string, chars string) int {
	cnts := toCharCount(chars)

	sum := 0
	for _, w := range words {
		wcnts := toCharCount(w)
		isGood := true
		for ch, cnt := range wcnts {
			if cnts[ch] < cnt {
				isGood = false
				break
			}
		}
		if isGood {
			sum += len(w)
		}
	}
	return sum
}

func toCharCount(s string) map[rune]int {
	m := map[rune]int{}
	for _, ch := range s {
		m[ch]++
	}
	return m
}
