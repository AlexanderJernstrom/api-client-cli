package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

func GetRequest(url string) (Response) {
	response, err := http.Get(url)
	var returnResponse Response

	if err != nil {
		panic(err)
	}

	var response_type string = response.Header["Content-Type"][0]
	var response_body interface{}

	returnResponse.ResponseType = response_type
	returnResponse.StatusCode = response.StatusCode

	if strings.Contains(response_type, "text/html") {
	} else if strings.Contains(response_type, "application/json") {
		if err := json.NewDecoder(response.Body).Decode(&response_body); err != nil {
			panic(err)
		}

		jsonBytes, err := json.Marshal(response_body)
		jsonString := string(jsonBytes)

		if err != nil {
			panic(err)
		}
		returnResponse.RawResponse = jsonString

	}

	return returnResponse
}