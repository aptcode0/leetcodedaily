package dec

// Problem: https://leetcode.com/problems/valid-anagram/description
/*
Input: s = "anagram", t = "nagaram"
Output: true
*/

func isAnagram(s string, t string) bool {
	freqs := findCharFreqs(s)
  
	n := len(freqs)
	for _, ch := range t {
	    freqs[ch]--
	    if freqs[ch] == 0 {
		  n--
	    } 
	    if freqs[ch] < 0 {
		  return false
	    }
	}
	return n == 0   
  }
  
  func findCharFreqs(s string) map[rune]int {
	freqs := make(map[rune]int)
  
	for _, ch := range s {
	    freqs[ch]++
	}
	
	return freqs
  }