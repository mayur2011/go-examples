package common

func GetDigitCount(num int) int {

	dc := 1
	for num > 9 {
		num = num / 10
		dc++
	}
	return dc
}
