package main

import 	"fmt"

func min(x int, y int) int{
	if x < y {
		return x
   }
	return y
}
func max(x int, y int) int{
	if x > y {
		return x
   }
	return y
}


func maxWater(arr []int, n int) int{
    // To store the maximum water
    waterTrapped := 0;
	left := arr[0]
	right := arr[n-1]
    // For every element of the array
    for i := 0; i < n; i++ {
 
        // Find the maximum element on its left
        left = arr[i];
        for j := 0; j < i; j++ {
            left = max(left, arr[j])
        }
 
        // Find the maximum element on its left
        right = arr[i];
        for j := i + 1; j < n; j++ {
            right = max(right, arr[j])
        }
 
        // Update the result (maximum water)
        waterTrapped = waterTrapped + (min(left, right) - arr[i])
    }
 
    // Return the maximum water trapped
    return waterTrapped;
}


func main() {
	arr := []int{4,2,0,3,2,5}
	size := len(arr)
	maxWaterTrapped := maxWater(arr, size)
	fmt.Printf("Trapped Water: %d\n", maxWaterTrapped)
}