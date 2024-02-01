package feb

import "sort"

// Problem: https://leetcode.com/problems/divide-array-into-arrays-with-max-difference/description
/*
Input: nums = [1,3,4,8,7,9,3,5,1], k = 2
Output: [[1,1,3],[3,4,5],[7,8,9]]
Explanation: We can divide the array into the following arrays: [1,1,3], [3,4,5] and [7,8,9].
The difference between any two elements in each array is less than or equal to 2.
Note that the order of elements is not important.
*/

func divideArray(nums []int, lim int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	
	var res [][]int
	for i := 0; i < n; i += 3 {
	    if nums[i+2] - nums[i] > lim {
		  return nil
	    }
	    res = append(res, nums[i:i+3])
	}
	return res
  }