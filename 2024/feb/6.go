package feb

// Problem: https://leetcode.com/problems/group-anagrams/description
/*
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
*/

import "sort"

func groupAnagrams(strs []string) [][]string {
	anagrams := map[string][]string{}
	for _, s := range strs {
	    sorted := sortString(s)
	    anagrams[sorted] = append(anagrams[sorted], s)
	}
	
	var res [][]string 
	for _, v := range anagrams {
	    res = append(res, v)
	}
	return res
  }
  
  func sortString(s string) string {
	chs := []byte(s)
	sort.Slice(chs, func(i, j int) bool { return chs[i] < chs[j] })
	return string(chs)
  }