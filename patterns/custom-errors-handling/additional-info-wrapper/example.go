package main

import (
	"fmt"
	"net/http"
)

type WrappedError struct {
	Context string
	Err     error
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.Context, w.Err)
}

func Wrap(err error, info string) *WrappedError {
	return &WrappedError{
		Context: info,
		Err:     err,
	}
}

func main() {
	res, err := http.Get("https://yaru")
	if err != nil {
		errResult := Wrap(err, "test")
		fmt.Println(*errResult)
	} else {
		fmt.Println(res.StatusCode)
	}
}
