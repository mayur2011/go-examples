package main

import (
	"fmt"
)

func main() {

	size := 5
	total := [5]int{2, 3, 1, 5, 4}
	//total := []int{0, 0, 0, 0, 0},	//total := []int{2, 3, 2, 5, 4},	//total := []int{3, 3, 1, 5, 4}
	index := 0
	y := 1
	smallest := total[0]
	fmt.Printf("Existing batch status :%d \n", total)
	for y <= size-1 {
		if total[y] < smallest {
			smallest = total[y]
			index = y
		}
		y++
	}
	total[index] = smallest + 1
	batchNum, assignedDay := getBatchNumAndWorkingday(index)
	printBatchInfo(batchNum, assignedDay, total)
}

func getBatchNumAndWorkingday(index int) (int, string) {
	batchNum := 0
	day := ""
	switch index {
	case 0:
		batchNum = 1
		day = "Monday" //ideally this needs to call a function which gives a day or date by considering weekends and company holidays

	case 1:
		batchNum = 2
		day = "Tuesday"

	case 2:
		batchNum = 3
		day = "Wednesday"

	case 3:
		batchNum = 4
		day = "Thursday"

	case 4:
		batchNum = 5
		day = "Friday"
	}
	return batchNum, day
}

func printBatchInfo(num int, day string, status interface{}) {
	fmt.Printf("Assigned Batch-Num : %d and Working-Day : %s\nBatch status : %d\n", num, day, status)
}
