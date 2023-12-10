package dec

// Problem: https://leetcode.com/problems/transpose-matrix/description
/*
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[1,4,7],[2,5,8],[3,6,9]]
*/

func transpose(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	
	res := make([][]int, cols)
	for i, _ := range res {
	    res[i] = make([]int, rows)
	}
  
	for i, _ := range matrix {
	    for j, num := range matrix[i] {
		  res[j][i] = num
	    }
	}
	
	return res
  }
  
  
