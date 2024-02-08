package feb

// Problem: https://leetcode.com/problems/perfect-squares/description
/*
Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
*/

func numSquares(n int) int {
	var sqrs []int
	
	for i := 1; i*i <= n; i++ {
	    sqrs = append(sqrs, i*i)    
	}
  
	q := []int{n}
	lvl := 1
	for len(q) > 0 {
	    for i := len(q); i > 0; i-- {
		  rem := q[0]
		  q = q[1:]
		  for _, s := range sqrs {
			if s > rem {
			    break
			}
			if s == rem {
			    return lvl
			}
			q = append(q, rem - s)
		  }
	    }
	    lvl++
	}
	return -1
  }