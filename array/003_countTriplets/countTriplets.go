package main

import (
	"fmt"
	"sort"
)

var p = fmt.Printf

func countTriplets(na []int, total int) (bool, []int) {
	sort.Ints(na)
	p("%d\n",na)

	for i:=0;i<len(na)-2;i++{
	  l, r := i+1, len(na)-1

	  for l < r {
	    curSum := na[i]+na[l]+na[r]
	    
	    if curSum == total {
		return true, []int{na[i],na[l],na[r]}
	    }else if curSum < total{
		l++
	    }else {
		r--
	    }	    
	  }
	}
	return false, []int{}
}

func main() {
	nums := []int{9, 1, 2, 4, 5, 7, 8}
	target := 9

	//size := len(nums)
	//p("size: %d\n", len(nums))
	c, res := countTriplets(nums, target)
	p("%v = %v\n", c, res)
}
