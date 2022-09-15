package api

import (
	"fmt"
	"net/http"
)

func GetRequest(url string) {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	var response_type string = response.Header["Content-Type"][0]

	fmt.Println(response_type)

	

}