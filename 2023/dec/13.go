package dec

// Problem : https://leetcode.com/problems/special-positions-in-a-binary-matrix/description
/*
Input: mat = [[1,0,0],[0,0,1],[1,0,0]]
Output: 1
Explanation: (1, 2) is a special position because mat[1][2] == 1 and all other elements in row 1 and column 2 are 0.
*/

func numSpecial(mat [][]int) int {
	cntInRows := map[int]int{}
	cntInCols := map[int]int{}
	var ones [][]int
	for i, r := range mat {
	    for j, num := range r {
		  if num == 1 {
			cntInRows[i]++
			cntInCols[j]++
			ones = append(ones, []int{i, j})
		  }
	    }
	}
	cnt := 0
	for _, pos := range ones {
	    if cntInRows[pos[0]] == 1 && cntInCols[pos[1]] == 1{
		  cnt++
	    }
	}    
	return cnt
  }