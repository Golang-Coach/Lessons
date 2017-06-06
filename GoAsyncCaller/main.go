package main

import (
	"fmt"
	"net/http"
	"time"
	"runtime"
)

type HTTPResponse struct {
	res http.Response
	err error
}

func main() {
	runtime.GOMAXPROCS(5)
	urls := []string{
		"https://httpbin.org/ip",
		"https://httpbin.org/get",
		"https://httpbin.org/gzip",
	}

	responses := asyncHttp(urls)
	for _, result := range responses {
		fmt.Printf("%+v \n",
			result)
	}
	fmt.Println(responses)
}

func asyncHttp(urls []string) []*HTTPResponse {
	ch := make(chan *HTTPResponse)
	responses := []*HTTPResponse{}

	for _, url := range urls {
		go func(url string) {
			fmt.Println(url)
			res, err := http.Get(url);
			res.Body.Close()
			ch <- &HTTPResponse{
				*res, err,
			}
		}(url)
	}

	for{
		select {
		case res := <-ch:
			fmt.Printf("%+v\n", res)
			responses = append(responses, res)

			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout")
			return responses;
		}
	}
	return responses
}
