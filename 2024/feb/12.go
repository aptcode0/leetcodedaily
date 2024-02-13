package feb

// Problem: https://leetcode.com/problems/majority-element/description
/*
Input: nums = [3,2,3]
Output: 3
*/

func majorityElement(nums []int) int {
	currMax := nums[0]
	k := 0
	for _, num := range nums {
		if num == currMax {
			k++
			continue
		}
		k--
		if k < 0 {
			currMax = num
			k = 1
		}
	}
	return currMax
}
