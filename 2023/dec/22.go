package dec

// Problem: https://leetcode.com/problems/maximum-score-after-splitting-a-string/description
/*
Input: s = "011101"
Output: 5 
Explanation: 
All possible ways of splitting s into two non-empty substrings are:
left = "0" and right = "11101", score = 1 + 4 = 5 
left = "01" and right = "1101", score = 1 + 3 = 4 
left = "011" and right = "101", score = 1 + 2 = 3 
left = "0111" and right = "01", score = 1 + 1 = 2 
left = "01110" and right = "1", score = 2 + 1 = 3
*/

func maxScore(s string) int {
	zeros := make([]int, len(s))
	for i := 1; i < len(s); i++ {
	    zeros[i] = zeros[i-1]
	    if s[i-1] == '0' {
		  zeros[i]++
	    }
	}    
	var max, ones int
	for i := len(s)-1; i > 0; i-- {
	    if s[i] == '1' {
		  ones++
	    }
	    curr := ones + zeros[i]
	    if curr > max {
		  max = curr
	    }
	}
	return max
  }