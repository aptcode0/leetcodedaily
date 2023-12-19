package dec

import "sort"

// Problem: https://leetcode.com/problems/maximum-product-difference-between-two-pairs/description
/*
Input: nums = [5,6,2,7,4]
Output: 34
Explanation: We can choose indices 1 and 3 for the first pair (6, 7) and indices 2 and 4 for the second pair (2, 4).
The product difference is (6 * 7) - (2 * 4) = 34.
*/

func maxProductDifference(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	return (nums[n-1] * nums[n-2]) - (nums[0] * nums[1])
}
