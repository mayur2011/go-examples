package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "Tjhis is from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server 2"
	}
}

func main() {
	fmt.Println("Select with channels")
	fmt.Println("--------------------")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	// endless for loop
	for {
		select {
		case s1 := <-channel1: // getting something from channel 1
			fmt.Println("Case one:", s1)
		case s2 := <-channel1:
			fmt.Println("Case two:", s2)
		case s3 := <-channel2:
			fmt.Println("Case three:", s3)
		case s4 := <-channel2:
			fmt.Println("Case four:", s4)

			//default:
			// useful for avoiding deadlock
			//fmt.Println("to Avoid deadlock")
		}
	}
}

/*
Note:
When this select statement executes it and there are multiple cases that match the same condition where we get responses back on for example line num: 39, 41
these are functinally identical. The only difference is that it's printing a slightly different message out here.
If there is more than one case that the select can match, it just chooses one at random and there are lots of situations where that's useful.
*/
