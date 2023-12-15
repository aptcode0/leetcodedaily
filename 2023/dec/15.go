package dec

// Problem: https://leetcode.com/problems/destination-city/description
/*
Input: paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
Output: "Sao Paulo" 
Explanation: Starting at "London" city you will reach "Sao Paulo" city which is the destination city. 
Your trip consist of: "London" -> "New York" -> "Lima" -> "Sao Paulo".
*/
func destCity(paths [][]string) string {
	out := map[string]bool{}
	for _, p := range paths {
	    out[p[0]] = true
	    if _, exist := out[p[1]]; !exist {
		  out[p[1]] = false
	    }
	} 
  
	for city, hasOutPath := range out {
	    if !hasOutPath {
		  return city
	    }
	} 
	return ""
  }