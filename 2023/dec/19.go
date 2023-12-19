package dec

// Problem: https://leetcode.com/problems/image-smoother/description
/*
Input: img = [[100,200,100],[200,50,200],[100,200,100]]
Output: [[137,141,137],[141,138,141],[137,141,137]]
*/

func imageSmoother(img [][]int) [][]int {
	res := make([][]int, len(img))
  
	for i, r := range img {
	    res[i] = make([]int, len(r))
	    for j, _ := range r {
		  res[i][j] = avg(img, i, j)
	    }
	}    
	return res
  }
  
  func avg(img [][]int, x, y int) int {
	sum := 0
	cnt := 0
	for i := -1; i <= 1; i++ {
	    for j := -1; j <= 1; j++ {
		  xx, yy := x+i, y+j
		  if isValid(xx, yy, len(img), len(img[0])) {
			sum += img[xx][yy]
			cnt++
		  }
	    }
	}
	return sum / cnt
  }
  
  func isValid(x, y, rows, cols int) bool {
	return x >= 0 && y >= 0 && x < rows && y < cols
  }