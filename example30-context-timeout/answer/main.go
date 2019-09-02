package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
  "net/http"
  "context"
  "time"
)


type response struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type callResponse struct {
	Resp *response
	Err  error
}

func helper(ctx context.Context) <-chan *callResponse {

	respChan := make(chan *callResponse, 1)

	go func() {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

		if err != nil {
			respChan <- &callResponse{nil, fmt.Errorf("error in http call")}
			return
		}

		defer resp.Body.Close()
		byteResp, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			respChan <- &callResponse{nil, fmt.Errorf("error in reading response")}
			return
		}

		structResp := &response{}
		err = json.Unmarshal(byteResp, structResp)

		if err != nil {
			respChan <- &callResponse{nil, fmt.Errorf("error in unmarshalling response")}
		}

		respChan <- &callResponse{structResp, nil}
	}()

	return respChan
}

func getHTTPResponse(ctx context.Context) (*response, error) {
  select {
	case <-ctx.Done():
		return nil, fmt.Errorf("context timeout, ran out of time")
	case respChan := <-helper(ctx):
		return respChan.Resp, respChan.Err

	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	res, err := getHTTPResponse(ctx)

	if err != nil {
		fmt.Printf("err %v", err)
	} else {
		fmt.Printf("res %v", res)
	}

}
