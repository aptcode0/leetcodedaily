package feb

// Problem: https://leetcode.com/problems/first-unique-character-in-a-string/description/
/*
Input: s = "leetcode"
Output: 0
*/

func firstUniqChar(s string) int {
    seen := map[byte]int{}
    
    for i := range s {
        seen[s[i]]++
    }
    
    for i := range s {
        if seen[s[i]] == 1 {
            return i
        }
    }
    return -1
}