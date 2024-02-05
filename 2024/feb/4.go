package main

// Problem: https://leetcode.com/problems/minimum-window-substring/description
/*
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
*/

func minWindow(s string, t string) string {
    freqs1 := calcFreqs(t)
    l := 0
    freqs2 := make(map[byte]int)
    matches := 0
    start, minLen := -1, len(s)+1
    for r, _ := range s {
        freqs2[s[r]]++
        if freqs1[s[r]] == freqs2[s[r]] {
            matches++
        }

        for matches == len(freqs1) {
            clen := r - l + 1
            if minLen > clen {
                minLen = clen
                start = l
            }
            if f, exist := freqs1[s[l]]; exist && f == freqs2[s[l]] {
                matches--
            }
            freqs2[s[l]]--
            l++
        }
    }
    if start == -1 {
        return ""
    }
    return s[start:start+minLen]
}

func calcFreqs(s string) map[byte]int {
    f := map[byte]int{}
    for i, _ := range s {
        f[s[i]]++
    }
    return f
}