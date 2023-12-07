package dec

// Problem: https://leetcode.com/problems/calculate-money-in-leetcode-bank/description
/*
Input: n = 4
Output: 10
Explanation: After the 4th day, the total is 1 + 2 + 3 + 4 = 10.
*/
func totalMoney(n int) int {
	fullWeekCount := n % 7
	remDays := n / 7

	// 1 + 2 + .. + 7 = 28
	// 2 + ...... + 8 = 35 = 28 + 7
	//  e.g. 21 = 3 fullweeks = 28+(28 + 7)+(28+7+7) = 3*28 + 7*1+7*2 = 3*28 + 7*(1+2)
	// sum formula = 1+2+3+..+n = n*(n+1)/2
	tot := remDays*28 + 7*((remDays-1)*remDays/2)
	tot += fullWeekCount*(fullWeekCount+1)/2 + remDays*fullWeekCount

	return tot
}
