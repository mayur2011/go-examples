package hello

import "fmt"

func SayHello(msg string, fn func() string) string {
	fmt.Println("Outside main method")
	s := fmt.Sprintf("Closure %s & ID: %s", msg, fn())
	return s
}
