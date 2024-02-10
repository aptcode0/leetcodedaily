package feb

// Problem: https://leetcode.com/problems/largest-divisible-subset/description
/*
Input: nums = [1,2,3]
Output: [1,2]
Explanation: [1,3] is also accepted.
*/
import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	parent := make([]int, len(nums))
	maxLen, maxPos := 0, 0

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		parent[i] = i
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] != 0 {
				continue
			}
			if dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				parent[i] = j
			}
		}
		if maxLen < dp[i] {
			maxLen = dp[i]
			maxPos = i
		}
	}

	curr := maxPos
	res := []int{nums[curr]}
	for parent[curr] != curr {
		res = append(res, nums[parent[curr]])
		curr = parent[curr]
	}
	return rev(res)
}

func rev(res []int) []int {
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
