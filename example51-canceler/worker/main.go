package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func cancelTask(id string) []byte {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/watch-task/"+id, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	return resBody
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			resp := string(cancelTask("1234"))
			fmt.Println("task[1234]: cancel the task:", resp)
			if resp == "true" {
				fmt.Println("task[1234]: get cancel event and canceld the task")
				return
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			resp := string(cancelTask("5678"))
			fmt.Println("task[5678]: cancel the task:", resp)
			if resp == "true" {
				fmt.Println("task[5678]: get cancel event and canceld the task")
				return
			}
		}
	}()
	wg.Wait()
}
