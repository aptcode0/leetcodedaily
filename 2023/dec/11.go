package dec

// Problem: https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/description
/*
Input: arr = [1,2,2,6,6,6,6,7,10]
Output: 6
*/
func findSpecialInteger(arr []int) int {
	nxtPosOffset := len(arr) / 4
	for i, num := range arr {
		pos := i + nxtPosOffset
		if pos >= len(arr) {
			break
		}
		if arr[pos] == num {
			return num
		}
	}
	return -1
}
