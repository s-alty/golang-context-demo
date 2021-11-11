package main

import (
	"context"
	"fmt"
	"reflect"
	"unsafe"
)

func getParent(ctx context.Context) context.Context{
	elem := reflect.ValueOf(ctx).Elem()
	parentField := elem.FieldByName("Context")
	return parentField.Interface().(context.Context)
}

func getUnexportedField(v reflect.Value, fieldName string) reflect.Value {
	// https://stackoverflow.com/questions/42664837/how-to-access-unexported-struct-fields
	unexportedField := v.FieldByName(fieldName)
	return reflect.NewAt(unexportedField.Type(), unsafe.Pointer(unexportedField.UnsafeAddr())).Elem()
}



// remove a child context from its parent's children map
// this saves the child from getting canceled when the parent is
func emancipate(child context.Context) {
	if child == context.Background() || child == context.TODO() {
		return
	}

	parent := getParent(child)

	// use reflection to access the parent's children field and remove the child from it
	elem := reflect.ValueOf(parent).Elem()
	childrenMap := getUnexportedField(elem, "children")

	// call SetMapIndex with the empty value (reflect.Value{}) to delete a key
	// make sure to only delete this child
	childVal := reflect.ValueOf(child)
	for _, k := range childrenMap.MapKeys(){
		if k.Elem() == childVal {
			childrenMap.SetMapIndex(k, reflect.Value{})
		}
	}

}


func main(){
	ctx, cancelFunc := context.WithCancel(context.Background())
	childCtx, _ := context.WithCancel(ctx)

	fmt.Println("Before")
	fmt.Printf("Parent status: %s\n", ctx.Err())
	fmt.Printf("Child status: %s\n", childCtx.Err())


	emancipate(childCtx)

	// normally cancelling the parent context would also cancel the child context
	cancelFunc()

	fmt.Println("****************************")
	fmt.Println("After")
	fmt.Printf("Parent status: %s\n", ctx.Err())
	fmt.Printf("Child status: %s\n", childCtx.Err())

}
