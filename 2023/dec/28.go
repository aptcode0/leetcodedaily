package dec

import "math"

// Problem: https://leetcode.com/problems/string-compression-ii/description
/*
Input: s = "aaabcccd", k = 2
Output: 4
*/
type state struct {
	idx int
	k int
	prev byte
	prevCnt int
  }

  var memo map[state]int
  
  func getLengthOfOptimalCompression(s string, k int) int {
	memo = map[state]int{}
	return findOptLen(s, state{0, k, 0, 0})
  }
  
  func findOptLen(s string, st state) int {
	if st.k < 0 {
	    return math.MaxInt - 1
	}
	if st.idx >= len(s) {
	    return 0
	}
	if v, exist := memo[st]; exist {
	    return v
	} 
  
	res := min(findDelLen(s, st), findKeepLen(s, st)) 
	memo[st] = res
	return res
  }
  
  func findDelLen(s string, st state) int {
	return findOptLen(s, state{st.idx+1, st.k-1, st.prev, st.prevCnt})
  }
  
  func findKeepLen(s string, st state) int {
	if s[st.idx] != st.prev {
	    return findOptLen(s, state{st.idx+1, st.k, s[st.idx], 1}) + 1
	}
	res := findOptLen(s, state{st.idx+1, st.k, st.prev, st.prevCnt+1})
	// s length max is 100 only
	if st.prevCnt == 1 || st.prevCnt == 9 || st.prevCnt == 99 {
	    res++
	}
	return res
  }

  func min(x, y int) int {
	if x < y {
		return x
	}
	return y
  }