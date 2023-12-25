package dec

// Problem: https://leetcode.com/problems/path-crossing/description
/*
Input: path = "NES"
Output: false 
Explanation: Notice that the path doesn't cross any point more than once.
*/

type point struct {
	x, y int
  }
  var dirs = map[byte][]int{
	'N': {-1, 0},
	'E': {0, 1},
	'S': {1, 0},
	'W': {0, -1},
  }
  func isPathCrossing(path string) bool {
	p := point{0, 0}
	seen := map[point]bool{}
	seen[p] = true
  
	for _, d := range []byte(path) {
	    np := point{p.x+dirs[d][0], p.y+dirs[d][1]}
	    if seen[np] {
		  return true
	    }
	    seen[np] = true
	    p = np
	}
	return false
  }