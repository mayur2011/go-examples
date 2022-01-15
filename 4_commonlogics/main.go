package main

import (
	"fmt"
	"go-examples/4_commonlogics/common"
)

func main() {
	data := []int{450, 654, 769, 10200, 9099, 110, 99}
	fmt.Println(data)
	x := data
	queue := [10][10]int{}
	rear := [10]int{}
	front := [10]int{}
	e := 1
	f := 10

	largest, _ := common.GetLargestNum(x)
	//fmt.Printf("Largest value is %d and at index of list[%d]\n", largest, index)
	dc := common.GetDigitCount(largest)
	//fmt.Printf("%d", dc)

	y := 0
	for y < dc {
		//logic to spread the data to respective queues
		i := 0
		for i < len(x) {
			num := (x[i] % f) / e
			queue[num][rear[num]] = x[i]
			rear[num]++
			i++
		}
		//fmt.Println(queue)

		//logic to collect back the data to an array
		i = 0
		j := 0
		for i <= 9 {
			for rear[i] != 0 {
				x[j] = queue[i][front[i]]
				front[i]++
				j++
				if front[i] == rear[i] {
					front[i] = 0
					rear[i] = 0
				}
			}
			i++
		}
		fmt.Println(x)
		e = e * 10
		f = f * 10
		y++
	}
}
