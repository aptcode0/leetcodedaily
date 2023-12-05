package dec

// Problem: https://leetcode.com/problems/largest-3-same-digit-number-in-string/description
/*
Input: num = "6777133339"
Output: "777"
Explanation: There are two distinct good integers: "777" and "333".
"777" is the largest, so we return "777".
*/
func largestGoodInteger(num string) string {
	max := -1
	l := 1
	res := ""
	for i := 1; i < len(num); i++ {
		if num[i] == num[i-1] {
			l++
		} else {
			l = 1
		}
		dig := int(num[i] - '0')
		if l == 3 && max < dig {
			max = dig
			res = num[i-2 : i+1]
		}
	}
	return res
}
