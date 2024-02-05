package feb

// Problem: https://leetcode.com/problems/partition-array-for-maximum-sum/description
/*
Input: arr = [1,15,7,9,2,5,10], k = 3
Output: 84
Explanation: arr becomes [15,15,15,9,10,10,10]
*/

func maxSumAfterPartitioning(arr []int, k int) int {
	memo := make([]int, len(arr))
	return maxSumPartition(arr, 0, k, memo)
}

func maxSumPartition(arr []int, i, k int, memo []int) int {
	if i >= len(arr) {
		return 0
	}
	if memo[i] > 0 {
		return memo[i]
	}

	maxNum := 0
	res := 0
	for j := 0; j < k && i+j < len(arr); j++ {
		maxNum = max(maxNum, arr[i+j])
		res = max(res, (j+1)*maxNum+maxSumPartition(arr, i+j+1, k, memo))
	}

	memo[i] = res
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}