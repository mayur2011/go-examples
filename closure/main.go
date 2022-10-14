package main

import (
	"fmt"
	"go-examples/closure/hello"
)

func getIDValue() string {
	return "ID-100-01"
}

func giveMeFunction() func() string {
	id := getIDValue()
	return func() string {
		fmt.Println("Closure is getting invoked")
		return id
	}
}

func main() {
	fmt.Println("Main starting with Closure Example")
	closureVar := giveMeFunction()
	recieve := hello.SayHello("MainMethod", closureVar)
	fmt.Println("Output: " + recieve)
}
