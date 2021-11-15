package main

import (
	"sync"
	"fmt"
	"time"

	"github.com/uber/tchannel-go"
	tchanjson "github.com/uber/tchannel-go/json"
)

func main(){
	serverAddr := "127.0.0.1:8081"
	tchan, _ := tchannel.NewChannel("client", nil)
	client := tchanjson.NewClient(tchan, "test-service", &tchanjson.ClientOptions{HostPort: serverAddr})

	// include a value and a deadline to see if they are transmitted at the protocol level at all
	ctx, cancelFunc := tchannel.NewContext(time.Second * 2)
	tchanCtx := tchanjson.WithHeaders(ctx, map[string]string{
		"example": "ABCDEFGHIJKL",
	})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err := client.Call(tchanCtx, "test-procedure", struct{}{}, nil)
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
