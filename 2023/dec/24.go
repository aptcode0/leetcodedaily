package dec

// Problem: https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/description

/*
Input: s = "0100"
Output: 1
Explanation: If you change the last character to '1', s will be "0101", which is alternating.
*/

func minOperations(s string) int {
	return min(findOp(s, '0'), findOp(s, '1'))
  }
  
  func findOp(s string, start byte) int {
	cnt := 0
	for i, _ := range s {
	    if (i % 2 == 0 && s[i] != start) || (i % 2 != 0 && s[i] == start) {
		  cnt++
	    }
	}
	return cnt
  }
  
  func min(x, y int) int {
	if x < y {
	    return x
	}
	return y
  }