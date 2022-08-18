package main

import (
	"fmt"
	"unsafe"
)

var p = fmt.Printf

func main() {
	pp := 10
	a := int32(123)
	b := int64(123)
	c := "foo"
	d := struct {
		FieldA float32
		FieldB string
	}{0, "bar"}

	fmt.Printf("pp: %T, %d\n", pp, unsafe.Sizeof(pp))
	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c))
	fmt.Printf("d: %T, %d\n", d, unsafe.Sizeof(d))

	arr := []int32{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	n := unsafe.Sizeof(arr) / unsafe.Sizeof(arr[0])
	p("arr: %T, %d\n", arr, unsafe.Sizeof(arr))
	p("arr[0]: %T, %d\n", arr[0], unsafe.Sizeof(arr[0]))
	p("n: %T, %d\n", n, unsafe.Sizeof(n))
	p("value of n is %d", n)
}
