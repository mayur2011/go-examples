package main

import (
	"fmt"
)

func trappedWater(height []int) int{
	n := len(height)-1
	maxLeft := height[0] //initializing max left as first index
	maxRight := height[n] //initializing max right as last index
	waterTrapped := 0 //0 as initial trapped water
	left := 1 // initial point as 1st element because first index (0) can't trap any water
	right := n-1 // 2nd last point as last element because last index (n) can't trap any water

	for left <= right {
		if maxLeft < maxRight{
			if height[left] >= maxLeft{
				maxLeft = height[left]
			}else{
				waterTrapped += maxLeft-height[left]
			}
			left++
		}else{
			if height[right]>= maxRight{
				maxRight = height[right]
			}else{
				waterTrapped += maxRight-height[right]
			}
			right--
		}
	}

	return waterTrapped
}


func main() {

	arr := []int{4,2,0,6,2,3,5}
	totalWaterTrapped := trappedWater(arr)
	fmt.Printf("Trapper Water: %d\n", totalWaterTrapped)
}
