package main

import (
	"fmt"
	"strings"
)

// Receive only channel "ping <-chan string"
// Send only channel "pong chan<- string"
func shout(ping <-chan string, pong chan<- string) {
	// run forever with no exit condition
	for {
		//get from channel ping and put to it variable s
		s, ok := <-ping

		// 2nd parameter (ok) -- just tells whether the receive value was sent on the channel (True) OR it's zero value return bcz the channel is closed & empty.
		if !ok {
			// do something
			fmt.Println(ok)
		}

		// converts it to uppercase, adds three exclamation (!!!) marks and sends that string back to the pong channel
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// create two channels
	ping := make(chan string)
	pong := make(chan string)

	// run func in the back-ground
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			// jump out this loop
			break
		}

		// send it to ping channel
		ping <- userInput

		// wait for a response
		response := <-pong
		fmt.Println("response:", response)
	}

	fmt.Println("All done, Closing channels.")
	close(ping)
	close(pong)
}
