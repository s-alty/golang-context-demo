package main

import (
	"context"
	"time"
	"fmt"
	"sync"

	"cross_network/grpc/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)




func main(){
	url := "127.0.0.1:8082"
	conn, _ := grpc.Dial(url, grpc.WithInsecure())
	defer conn.Close()
	client := service.NewNoopClient(conn)

	// include a value and a deadline to see if they are transmitted at the protocol level at all
	ctx, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Second * 2,
	)

	d := metadata.New(map[string]string{
		"example": "ABCDEFGHIJKL",
	})
	ctx = metadata.NewOutgoingContext(ctx, d)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		_, err := client.Noop(ctx, &service.NoopRequest{})
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
