package feb

// Problem: https://leetcode.com/problems/cherry-pickup-ii/description
/*
Input: grid = [[3,1,1],[2,5,1],[1,5,5],[2,1,1]]
Output: 24
Explanation: Path of robot #1 and #2 are described in color green and blue respectively.
Cherries taken by Robot #1, (3 + 2 + 5 + 2) = 12.
Cherries taken by Robot #2, (1 + 5 + 5 + 1) = 12.
Total of cherries: 12 + 12 = 24.
*/

type state struct {
	r  int
	c1 int
	c2 int
}

func cherryPickup(grid [][]int) int {
	memo := map[state]int{}
	return pickCherry(grid, 0, 0, len(grid[0])-1, memo)
}

func pickCherry(grid [][]int, r, c1, c2 int, memo map[state]int) int {
	if r == len(grid) || c1 < 0 || c1 >= len(grid[0]) || c2 < 0 || c2 >= len(grid[0]) {
		return 0
	}

	state := state{r, c1, c2}
	if v, exist := memo[state]; exist {
		return v
	}

	maxSum := 0
	for i1 := -1; i1 <= 1; i1++ {
		for i2 := -1; i2 <= 1; i2++ {
			maxSum = max(maxSum, pickCherry(grid, r+1, c1+i1, c2+i2, memo))
		}
	}

	maxSum += grid[r][c1]
	if c1 != c2 {
		maxSum += grid[r][c2]
	}
	memo[state] = maxSum
	return maxSum
}

func getCherry(grid [][]int, r, c int) int {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
		return 0
	}
	return grid[r][c]
}
