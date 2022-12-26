package main

import "fmt"

func main() {
 var a int
	fmt.Scanf("%d",&a)

sortedList := []int{3,5,6,8,9,10,13,17,19,22,24,31}
inputValue := a
lb := 0
ub := len(sortedList)-1

index := binarySearchSort(sortedList, inputValue, lb, ub)
if index >= 0 {
	fmt.Println("Value:",inputValue," is found at Index:",index)
} else {
	fmt.Println("Did not find the index for Value:",inputValue)
}

}

func binarySearchSort(list []int, value, lb, ub int) int {

	lo := lb
	hi := ub
	for lo <= hi {

	mid := lo+(hi-lo)/2
	midValue := list[mid]

	if midValue == value {
		//found the index
		return mid
	} else if value < midValue {
		//use the left half of list
		hi = mid-1
	} else {
		//use the right hafl of the list
		lo = mid+1
	}
}

	return -1
}
