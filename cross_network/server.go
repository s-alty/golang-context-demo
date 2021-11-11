package main

import (
	"net/http"
	"fmt"
)


// see if the value was present on the request context
// see if the deadline was set on the request context
// just do nothing until the context gets canceled
func handleRequest(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	value := ctx.Value("example")
	fmt.Printf("Value retrieved from context for key 'example': %v\n", value)
	_, ok := ctx.Deadline()
	fmt.Printf("A deadline was present on the context?: %v\n", ok)

	<- ctx.Done() // wait for cancellation
	fmt.Println("request was cancelled")
}

func main(){
	s := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: http.HandlerFunc(handleRequest),
		IdleTimeout: 0,
		ReadTimeout: 0,
	}
	s.ListenAndServe()
}
