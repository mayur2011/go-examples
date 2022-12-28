package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type Backend struct {
	net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
}

var backendQueue chan *Backend

func init() {
	backendQueue = make(chan *Backend, 8)
}

func main() {
	fmt.Println("starting main()")
	requests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	for _, v := range requests {
		go handleConnection(fmt.Sprintf("%s-%d", "Request-", v))
	}
	/*time.Sleep(2 * time.Second)
	handleConnection("connect-002------------------>")
	*/
	//time.Sleep(10 * time.Second)
	fmt.Println("ending main()")

}

// handleConnection is spawned once per connection from a client, and exits when the client is
// done sending requests.
func handleConnection(conn string) {
	fmt.Println("Processing", conn)
	be, err := getBackend()
	if err != nil {
		return
	}

	go queueBackend(be)
}

// getBackend either gets a backend from the queue or makes a new backend.
func getBackend() (*Backend, error) {
	//fmt.Println("fn:getBackend")
	select {
	case be := <-backendQueue:
		fmt.Println("Getting Conn from backendQueue")
		return be, nil
	case <-time.After(100 * time.Millisecond):
		//be, err := net.Dial("tcp", "127.0.0.1:8081")
		be, err := netDial("tcp", "127.0.0.1:8081")
		if err != nil {
			return nil, err
		}
		return &Backend{
			Conn:   be,
			Reader: bufio.NewReader(be),
			Writer: bufio.NewWriter(be),
		}, nil
	}
}

// queueBackend takes a backend and reenqueues it.
func queueBackend(be *Backend) {
	select {
	case backendQueue <- be:
		fmt.Println("Putting to Queue")
		// Backend re-enqueued safely, move on.
	case <-time.After(50 * time.Millisecond):
		// If this fires, it means the queue of backends was full. We don't want to wait
		// forever, as this period of time blocks us handling the next request a user
		// might send us.
		netClose()
		//be.Close()
	}
}

func netDial(port, addr string) (net.Conn, error) {
	fmt.Println("Conn not available, Creating it with netDial..!")
	return nil, nil
}

func netClose() {
	fmt.Println("Closing Connection..!")
}
