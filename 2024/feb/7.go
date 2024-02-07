package feb

// Problem: https://leetcode.com/problems/sort-characters-by-frequency/description
/*
Input: s = "tree"
Output: "eert"
Explanation: 'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
*/

import "sort"

type CharFreq struct {
	ch byte
	freq int
  }
  
  func frequencySort(s string) string {
	pairs := createFreqCharPairs(s)
	sortByFreq(pairs)
	return buildSortedString(pairs)
  }
  
  func createFreqCharPairs(s string) []CharFreq {
	freqs := createFreqMap(s)
	return createPairs(freqs)
  }
  
  func createFreqMap(s string) map[byte]int {
	freqs := map[byte]int{}
	for i := range s {
	    freqs[s[i]]++
	}
	return freqs
  }
  
  func createPairs(freqs map[byte]int) []CharFreq {
	var res []CharFreq
	for ch, freq := range freqs {
	    res = append(res, CharFreq{ch, freq})
	}
	return res
  }
  
  func sortByFreq(pairs []CharFreq) {
	sort.Slice(pairs, func(i, j int) bool { 
	    return pairs[i].freq > pairs[j].freq
	})
  }
  
  func buildSortedString(pairs []CharFreq) string {
	var chs []byte
	for _, p := range pairs {
	    for i := 0; i < p.freq; i++ {
		  chs = append(chs, p.ch)
	    }
	}
	return string(chs)
  }