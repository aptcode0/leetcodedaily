package dec

// Problem: https://leetcode.com/problems/buy-two-chocolates/description
/*
Input: prices = [1,2,2], money = 3
Output: 0
Explanation: Purchase the chocolates priced at 1 and 2 units respectively. You will have 3 - 3 = 0 units of money afterwards. 
Thus, we return 0.
*/

func buyChoco(prices []int, money int) int {
	min1, min2 := min(prices[0], prices[1]), max(prices[0], prices[1])
	for i := 2; i < len(prices); i++ {
	    p := prices[i]
	    if min1 > p {
		  min2 = min1
		  min1 = p
		  continue
	    }
	    if min2 > p {
		  min2 = p
	    }
	}
	sum := min1 + min2 
	if sum <= money {
	    return money - sum
	}
	return money
  }
  func min(x, y int) int {
	if x < y {
	    return x
	}
	return y
  }
  func max(x, y int) int {
	if x > y {
	    return x
	}
	return y
  }