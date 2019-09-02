package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type Response struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetHttpResponse() (*Response, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		return nil, fmt.Errorf("error in http call")
	}

	defer resp.Body.Close()
	byteResp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("error in reading response")
	}

	structResp := &Response{}
	err = json.Unmarshal(byteResp, structResp)

	if err != nil {
		return nil, fmt.Errorf("error in unmarshalling response")
	}

	return structResp, nil
}

func main() {

	res, err := GetHttpResponse()

	if err != nil {
		fmt.Printf("err %v", err)
	} else {
		fmt.Printf("res %v", res)
	}

}
