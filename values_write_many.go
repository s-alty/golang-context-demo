package main

import (
	"context"
	"fmt"
)

func SetMany(ctx context.Context, keys, values []interface{}) context.Context {
	if len(keys) == 0 {
		return ctx
	}
	updated := context.WithValue(ctx, keys[0], values[0])
	return SetMany(updated, keys[1:], values[1:])
}

func main(){
	result := SetMany(
		context.Background(),
		[]interface{}{"make", "model", "cc"},
		[]interface{}{"kawaksaki", "klx", 250},
	)
	make := result.Value("make")
	model := result.Value("model")
	cc := result.Value("cc")

	fmt.Printf("%s %s: %dcc\n", make, model, cc)
}
