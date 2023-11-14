package main

import (
	"fmt"
	"sort"
)

var p = fmt.Printf

func countTriplets(na []int, total int) bool {
	sort.Ints(na)

	for i:=0;i<len(na)-2;i++{
	  l, r := i+1, len(na)-1

	  for l < r {
	    curSum := na[i]+na[l]+na[r]
	    
	    if curSum == total {
		return true
	    }else if curSum < total{
		l++
	    }else {
		r--
	    }	    
	  }
	}
	return false
}

func main() {
	nums := []int{-3, 1, 2, 4, 5, 7, 8}
	target := 10

	size := len(nums)
	p("size: %d\n", len(arr))

	sort.Ints(arr)
	p("%d", arr)
	c := countTriplets(arr, target)
	p("%v", c)
}
