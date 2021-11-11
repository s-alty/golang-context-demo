package main

import (
	"context"
	"net/http"
	"time"
	"fmt"
	"sync"
)




func main(){
	url := "http://127.0.0.1:8080/"
	client := &http.Client{}

	// include a value and a deadline to see if they are transmitted at the protocol level at all
	ctx, cancelFunc := context.WithTimeout(
		context.WithValue(context.Background(), "example", "ABCDEFGHIJKL"),
		time.Second * 2,
	)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		_, err := client.Do(req)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
		wg.Done()
	}()

	time.Sleep(time.Second)
	// how does cancellation manifest at the network level?
	cancelFunc()

	wg.Wait() // make sure we give the goroutine enough time to exit
}
