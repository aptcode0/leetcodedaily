package dec

import (
	"container/heap"
	"strings"
)

// Problem: https://leetcode.com/problems/design-a-food-rating-system/description
/*
Input
["FoodRatings", "highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"]
[[["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], ["korean", "japanese", "japanese", "greek", "japanese", "korean"], [9, 12, 8, 15, 14, 7]], ["korean"], ["japanese"], ["sushi", 16], ["japanese"], ["ramen", 16], ["japanese"]]
Output
[null, "kimchi", "ramen", null, "sushi", null, "ramen"]
*/

type FoodRatings struct {
	fbyc map[string]*heapElem
	ref map[string]*data
  }
  
  
  func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fbyc := map[string]*heapElem{}
	ref := map[string]*data{}
	for i, _ := range foods {
	    if fbyc[cuisines[i]] == nil {
		  fbyc[cuisines[i]] = &heapElem{}
	    }
	    d := data{foods[i], ratings[i], cuisines[i]}
	    heap.Push(fbyc[cuisines[i]], &d)
	    ref[foods[i]] = &d
	}    
	return FoodRatings{fbyc, ref}
  }
  
  
  func (this *FoodRatings) ChangeRating(food string, newRating int)  {
	d := this.ref[food]
	d.rating = newRating
	heap.Init(this.fbyc[d.cuisine])
  }
  
  func (this *FoodRatings) HighestRated(cuisine string) string {
	return this.fbyc[cuisine].Top().food
  }
  
  type data struct {
	food string
	rating int
	cuisine string
  }
  
  type heapElem []*data 
  
  func (h heapElem) Len() int { return len(h) }
  func (h heapElem) Less(i, j int) bool { 
	if h[i].rating != h[j].rating {
	    return h[i].rating > h[j].rating
	}
	return strings.Compare(h[i].food, h[j].food) < 0
  }
  
  func (h heapElem) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
  
  func (h *heapElem) Push(v any) {
	*h = append(*h, v.(*data))
  }
  
  func (h *heapElem) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
  }
  func (h heapElem) Top() *data {
	return h[0]
  }
  
  /**
   * Your FoodRatings object will be instantiated and called as such:
   * obj := Constructor(foods, cuisines, ratings);
   * obj.ChangeRating(food,newRating);
   * param_2 := obj.HighestRated(cuisine);
   */