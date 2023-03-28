package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func Request(url string) RetrayableFunc {
	return func() error {
		resp, err := http.Get(url)
		if err != nil {
			return errors.New(fmt.Sprintf("Error getting HTTP response: %v", err))
		}
		fmt.Println(resp.StatusCode)
		return nil
	}
}

type RetrayableFunc func() error

func RetryExecution(fn RetrayableFunc, retries int, delay time.Duration) (err error) {
	for retry := 0; retry < retries; retry++ {
		err = fn()
		if err == nil {
			break
		}
		fmt.Printf("Failed to execute #%d", retry)
		time.Sleep(delay)
	}
	return err
}

func main() {
	err := RetryExecution(Request("https://examplecom"), 3, 10*time.Second)
	if err != nil {
		fmt.Printf("Failed to execute %d", err)
	}
}
