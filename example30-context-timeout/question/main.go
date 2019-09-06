package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getHTTPResponse() (*response, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		return nil, fmt.Errorf("error in http call")
	}

	defer resp.Body.Close()
	byteResp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("error in reading response")
	}

	structResp := &response{}
	err = json.Unmarshal(byteResp, structResp)

	if err != nil {
		return nil, fmt.Errorf("error in unmarshalling response")
	}

	return structResp, nil
}

func main() {

	res, err := getHTTPResponse()

	if err != nil {
		fmt.Printf("err %v", err)
	} else {
		fmt.Printf("res %v", res)
	}
}
