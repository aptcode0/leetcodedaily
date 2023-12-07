package dec

//Problem: https://leetcode.com/problems/largest-odd-number-in-string/description
/*
Input: num = "52"
Output: "5"
Explanation: The only non-empty substrings are "5", "2", and "52". "5" is the only odd number.
*/

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		dig := int(num[i] - '0')
		if dig%2 != 0 {
			return num[:i+1]
		}
	}
	return ""
}
