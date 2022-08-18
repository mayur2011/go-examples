package main

import "fmt"

//min function to return small int value
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//maxArea function to calculate container area
func maxArea(height []int) int {
	//setting left height with i
	i := 0
	//setting right height with j
	j := len(height) - 1

	//initializing 0 as area value
	maxAreaValue := 0

	//looping through i to j index (left to right iteration)
	for i < j {

		//getting min height value using customized min func
		minHeight := min(height[i], height[j])
		// setting width value
		width := j - i
		//calculating area of a container
		area := minHeight * width
		//checking the max area possible value
		if area > maxAreaValue {
			maxAreaValue = area
		}

		//left or right height wise moving the index
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxAreaValue
}

func main() {

	//initializing int array
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxAreaValue := maxArea(arr)
	fmt.Println(maxAreaValue)
}

/*
This code is contributed by Mayur Upadhyay
Go playground link -- https://go.dev/play/p/wQ2JLPcwNZc
*/
