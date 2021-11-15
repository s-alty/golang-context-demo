package main

import (
	"fmt"
	"context"
	"time"

	"github.com/uber/tchannel-go"
	tchanjson "github.com/uber/tchannel-go/json"
)

type Payload struct {}
type Response struct {}

// see if the value was present on the request context
// see if the deadline was set on the request context
// just do nothing until the context gets canceled
func handleRequest(ctx tchanjson.Context, p *Payload) (*Response, error){
	value := ctx.Headers()["example"]
	fmt.Printf("Value retrieved from headers for key 'example': %v\n", value)
	deadline, ok := ctx.Deadline()
	fmt.Printf("A deadline was present on the context?: %v\n", ok)
	fmt.Printf("Deadline is in %v\n", time.Until(deadline))

	<- ctx.Done() // wait for cancellation
	fmt.Println("request was cancelled")
	return &Response{}, nil
}

func main(){
	tchan, _ := tchannel.NewChannel("test-service", nil)
	tchanjson.Register(
		tchan,
		tchanjson.Handlers{
			"test-procedure": handleRequest,
		},
		func (ctx context.Context, err error){ return },
	)
	tchan.ListenAndServe("127.0.0.1:8081")
	select {}
}
