package dec
// Problem: https://leetcode.com/problems/minimum-time-visiting-all-points/description
/*
Input: points = [[1,1],[3,4],[-1,0]]
Output: 7
Explanation: One optimal path is [1,1] -> [2,2] -> [3,3] -> [3,4] -> [2,3] -> [1,2] -> [0,1] -> [-1,0]   
Time from [1,1] to [3,4] = 3 seconds 
Time from [3,4] to [-1,0] = 4 seconds
Total time = 7 seconds
*/

func minTimeToVisitAllPoints(points [][]int) int {
	sum := 0
	for i := 1; i < len(points); i++ {
	  sum += max(abs(points[i][0] - points[i-1][0]), abs(points[i][1] - points[i-1][1]))
	}    
	return sum
    }
    
    func abs(x int) int {
	if x < 0 {
	  return -x
	}
	return x
    }
    
    func max(x, y int) int {
	if x > y {
	  return x
	}
	return y
    }
    