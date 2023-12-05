package dec
// Problem: https://leetcode.com/problems/count-of-matches-in-tournament/description
/*
Input: n = 7
Output: 6
Explanation: Details of the tournament: 
- 1st Round: Teams = 7, Matches = 3, and 4 teams advance.
- 2nd Round: Teams = 4, Matches = 2, and 2 teams advance.
- 3rd Round: Teams = 2, Matches = 1, and 1 team is declared the winner.
Total number of matches = 3 + 2 + 1 = 6.
*/

func numberOfMatches(n int) int {
	sum := 0
	for n > 1 {
	  sum += n / 2
	  n -= n / 2
	}
	return sum
}

/*
OR simply
func numberOfMatches(n int) int {
	return n-1
}
*/
