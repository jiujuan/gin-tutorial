package main

import (
	"net/http"
	"time"

	"fmt"

	"github.com/afex/hystrix-go/hystrix"
)

type Hello struct{}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Command(w, r)
}

func (h *Hello) Command(w http.ResponseWriter, r *http.Request) {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:                int(time.Second * 3),
		MaxConcurrentRequests:  10,
		SleepWindow:            200,
		RequestVolumeThreshold: 10,
		ErrorPercentThreshold:  20,
	})

	msg := "ret success"
	_ = hystrix.Do("my_command",
		func() error {
			_, err := http.Get("http://httpbin.org")
			if err != nil {
				fmt.Println("http get error： ", err)
				return err
			}
			return nil
		},
		func(err error) error { // fallback 处理
			fmt.Printf("handler error： %v\n", err)
			msg = "ret error"
			return nil
		},
	)
	w.Write([]byte(msg))
}

func main() {
	http.ListenAndServe(":8080", &Hello{})
}
