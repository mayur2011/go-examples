package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	httpUrl := "https://api.github.com"

	client := http.Client{}
	response, err := client.Get(httpUrl)

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
this http.Get only get url and initiate a call to given url without configuring the headers
*/
