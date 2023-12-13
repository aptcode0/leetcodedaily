package dec

// Problem: https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/description
/*
Input: nums = [3,4,5,2]
Output: 12 
Explanation: If you choose the indices i=1 and j=2 (indexed from 0), you will get the maximum value, 
that is, (nums[1]-1)*(nums[2]-1) = (4-1)*(5-1) = 3*4 = 12. 
*/

func maxProduct(nums []int) int {
	max1, max2 := max(nums[0], nums[1]), min(nums[0], nums[1])
	
	for i := 2; i < len(nums); i++ {
	    num := nums[i]
	    if num > max1 {
		  max2 = max1
		  max1 = num
		  continue
	    }
	    if num > max2 {
		  max2 = num
	    }
	}
	return (max1-1) * (max2-1)
  }
  
  func max(x, y int) int {
	if x > y {
	    return x
	}
	return y
  }
  
  func min(x, y int) int {
	if x < y {
	    return x
	}
	return y
  }