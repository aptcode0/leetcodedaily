package dec

// Problem: https://leetcode.com/problems/minimum-time-to-make-rope-colorful/description
/*
Input: colors = "abaac", neededTime = [1,2,3,4,5]
Output: 3
Explanation: In the above image, 'a' is blue, 'b' is red, and 'c' is green.
Bob can remove the blue balloon at index 2. This takes 3 seconds.
There are no longer two consecutive balloons of the same color. Total time = 3.
*/
func minCost(colors string, neededTime []int) int {
	curr, res := 0, 0

	for i, t := range neededTime {
		if i > 0 && colors[i] != colors[i-1] {
			curr = 0
		}
		res += min(curr, t)
		curr = max(curr, t)
	}
	return res
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
