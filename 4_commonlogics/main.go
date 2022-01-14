package main

import (
	"fmt"
	"go-examples/4_commonlogics/common"
)

func main() {
	data := []int{45, 654, 769, 1200, 110}
	fmt.Println(data)
	queue := [10][4]int{}
	rear := [10]int{}
	//front := [10]int{}
	//x := [10]int{}
	e := 1
	f := 10

	largest, _ := common.GetLargestNum(data)
	//fmt.Printf("Largest value is %d and at index of list[%d]\n", largest, index)
	dc := common.GetDigitCount(largest)
	//fmt.Printf("%d", dc)

	y := 0
	for y < dc {

		//logic to spread the data to respective queues
		i := 0
		for i < len(data) {
			num := (data[i] % f) / e
			queue[num][rear[num]] = data[i]
			rear[num]++
			i++
		}

		//logic to collect back the data to an array

		// i := 0
		// for i < len(data)

		fmt.Println(queue)
		e = e * 10
		f = f * 10
		y++
	}
}
