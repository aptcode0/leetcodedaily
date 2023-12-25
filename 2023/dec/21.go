package dec

import "sort"

// Problem: https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/description
/*
Input: points = [[8,7],[9,9],[7,4],[9,7]]
Output: 1
Explanation: Both the red and the blue area are optimal.
*/

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func (i, j int) bool { return points[i][0] < points[j][0] })
  
	max := 0
	for i := 0; i < len(points)-1; i++ {
	    width := points[i+1][0]-points[i][0]
	    if max < width {
		  max = width
	    }
	}
	return max
  }