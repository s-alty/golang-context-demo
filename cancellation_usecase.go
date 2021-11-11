package main

import (
	"context"
	"fmt"
	"net"
	"time"
)



// return the first port we were able to connect to
func portScan(ip string) (int, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	resultChan := make(chan int)
	for port := 1; port < 1024; port++ {
		go testConnection(ctx, ip, port, resultChan)
	}


	// as soon as we get one result back cancel the rest of them to avoid unneeded work
	defer cancelFunc()

	timeLimit := time.After(time.Second * 5)
	select {
	case result := <- resultChan:
		return result, nil
	case <- timeLimit:
		return 0, fmt.Errorf("Timed out before finding any open ports")
	}
}

func testConnection(ctx context.Context, ip string, port int, resultChan chan<- int){
	// internally net.DialContext will respect context cancellation
	var d net.Dialer
	_, err := d.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return
	}

	// here we use select with ctx.Done() to avoid blocking on the channel write if we've already obtained a result
	select {
		case <-ctx.Done():
		case resultChan <- port:
	}
}


func main(){
	ip := "127.0.0.1"
	port, err := portScan(ip)
	if err == nil {
		fmt.Printf("Found listener at %s:%d\n", ip, port)
	} else {
		fmt.Printf(err.Error())
	}
}
