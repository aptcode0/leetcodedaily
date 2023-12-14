package dec

// Problem: https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column/description
/*
Input: grid = [[0,1,1],[1,0,1],[0,0,1]]
Output: [[0,0,4],[0,0,4],[-2,-2,2]]
*/

func onesMinusZeros(grid [][]int) [][]int {
	rows, cols := len(grid), len(grid[0])
	cntRows := make([][]int, 2)
	for i, _ := range cntRows {
	    cntRows[i] = make([]int, rows)
	}
	cntCols := make([][]int, 2)
	for i, _ := range cntCols {
	    cntCols[i] = make([]int, cols)
	}
  
	for i, r := range grid {
	    for j, v := range r {
		  cntRows[v][i]++
		  cntCols[v][j]++
	    }
	}
  
	res := make([][]int, rows)
	for i, _ := range res {
	    res[i] = make([]int, cols) 
	    for j, _ := range res[i] {
		  res[i][j] = cntRows[1][i]+cntCols[1][j] -
		  cntRows[0][i] - cntCols[0][j]
	    }
	}
	return res
  }