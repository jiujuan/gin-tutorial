package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    200,
			IdleConnTimeout: 2 * time.Second,
		},
	}

	var wg sync.WaitGroup

	for i := 0; i < 13; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			req, err := http.NewRequest("GET", "http://localhost:8080", nil)
			if err != nil {
				fmt.Println("init http client error: ", err)
				return
			}
			rsp, err := client.Do(req)
			if err != nil {
				fmt.Println("do http error: ", err)
				return
			}

			defer rsp.Body.Close()

			res, err := ioutil.ReadAll(rsp.Body)
			if err != nil {
				fmt.Println("read body error: ", err)
				return
			}
			fmt.Println("read body: ", string(res))
		}()
	}

	wg.Wait()

	fmt.Println("finished")
}
