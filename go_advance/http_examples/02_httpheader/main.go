package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	httpMethod := "GET"
	url := "https://api.github.com"

	client := http.Client{}

	request, err := http.NewRequest(httpMethod, url, nil)
	//request.Header.Set("Accept", "application/xml")
	request.Header.Set("Accept", "application/json")
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}

/*
- The headers are part of the request, not of client
- Every request that you have will have different headers
- We need some way of creating our own create and start adding headers to the request
- http.NewRequest - can be used to create our own request
- client.Do instead of Client.Get needs to be used to pass our request

Question: So if it is that easy, why we are creating our own client?
- refer next program
*/
