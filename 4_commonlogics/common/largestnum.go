package common

func GetLargestNum(x []int) (int, int) {
	//fmt.Printf("%v \n", x)
	length := len(x)
	//fmt.Printf("%d", length)

	largest := x[0]
	index := 0
	y := 0
	for y < length {
		if x[y] > largest {
			index = y
			largest = x[y]
		}
		y++
	}
	return largest, index
}
