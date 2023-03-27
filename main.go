package main

import (
	"context"
	"fmt"
	"time"

	"github.com/yukihir0/expbo/expbo"
)

var requestCount = 0

// リクエスト3回目まではエラーを返す関数
func request() error {
	requestCount = requestCount + 1

	if requestCount > 3 {
		return nil
	} else {
		return fmt.Errorf("error occured.")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	maxRetryCount := 5
	e := expbo.NewExponentialBackoffJitter(
		ctx,
		maxRetryCount,
	)

	for t := range e.Retry() {
		err := request()
		if err != nil {
			fmt.Printf("retry... sleep: %d ms\n", t)
			time.Sleep(t * time.Millisecond)
			continue
		}

		fmt.Println("success")
		cancel()
	}
}
