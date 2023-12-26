package dec

// Problem: https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/description
/*
Input: n = 1, k = 6, target = 3
Output: 1
Explanation: You throw one die with 6 faces.
There is only one way to get a sum of 3.
*/

func numRollsToTarget(n int, k int, target int) int {
	dp := initDp(n, target)
	dp[0][0] = 1
	const mod int = 1e9 + 7
	for i := 1; i <= n; i++ {
	    for t := 1; t <= target; t++ {
		  for j := 1; j <= k && j <= t; j++ {
			dp[i][t] = (dp[i][t] + dp[i-1][t-j]) % mod
		  }
	    }
	}
	return dp[n][target] 
  }
  
  func initDp(n, t int) [][]int {
	dp := make([][]int, n+1)
	for i, _ := range dp {
	    dp[i] = make([]int, t+1)
	}
	return dp
  }
  
  // Recursive solution, memoization can be added
  func dp(n, k, t int) int {
	if n == 0 && t == 0 {
	    return 1
	}
	if n == 0 || t == 0 {
	    return 0
	}
  
	ans := 0
	for j := 1; j <= k; j++ {
	    ans += dp(n-1, k, t-j)
	} 
  
	return ans
  }