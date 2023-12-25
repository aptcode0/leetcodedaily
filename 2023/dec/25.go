package dec

// Problem: https://leetcode.com/problems/decode-ways/description
/*
Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
*/

func numDecodings(s string) int {
	cache := make([]int, len(s)+1)
	cache[len(s)] = 1
	for i := len(s)-1; i >= 0; i-- {
	    if s[i] != '0' {
		  cache[i] += cache[i+1]
	    }
	    if i+1 < len(s) && isValid(s[i], s[i+1]) {
		  cache[i] += cache[i+2]
	    }
	}
	return cache[0]
  }
  
  func isValid(n1, n2 byte) bool {
	if n1 == '1' {
	    return true
	}
	if n1 == '2' {
	    return n2 <= '6'
	}
	return false
  }