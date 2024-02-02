package feb

// Problem: https://leetcode.com/problems/sequential-digits/description
/*
Input: low = 100, high = 300
Output: [123,234]
*/

func sequentialDigits(low int, high int) []int {
	end := 9
	lowLen, highLen := findLen(low), findLen(high)
	var res []int
  
	for l := lowLen; l <= highLen; l++ {
	    for start := 1; start <= end-l+1; start++ {
		  num := generateSeq(start, l)
		  if num >= low && num <= high {
			res = append(res, num)
		  }
	    }
	}
	
	return res
  }
  
  func findLen(n int) int {
	cnt := 0
	for n > 0 {
	    n /= 10
	    cnt++
	}
	return cnt
  }
  
  func generateSeq(start, size int) int {
	curr := 0
	for i := 0; i < size; i++ {
	    curr = curr * 10 + (start + i)
	}
	return curr
  }
  
  