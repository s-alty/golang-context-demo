package main

import
(
	"context"
	"reflect"
	"unsafe"
	"fmt"
)


func GetAllValues(ctx context.Context) ([]interface{}, []interface{}) {
	// Base case: the root contexts have no values
	if ctx == context.Background() || ctx == context.TODO() {
		return []interface{}{}, []interface{}{}
	}


	// use reflection to get to the parent contex
	elem := reflect.ValueOf(ctx).Elem()
	parentField := elem.FieldByName("Context")
	parent := parentField.Interface().(context.Context)

	// recurse
	parentKeys, parentValues := GetAllValues(parent)

	// lookup the current key and value using reflection
	currentKey := getUnexportedField(elem, "key")
	currentValue := getUnexportedField(elem, "val")

	return append(parentKeys, currentKey), append(parentValues, currentValue)
}

func getUnexportedField(v reflect.Value, fieldName string) interface{} {
	// https://stackoverflow.com/questions/42664837/how-to-access-unexported-struct-fields
	unexportedField := v.FieldByName(fieldName)
	exported := reflect.NewAt(unexportedField.Type(), unsafe.Pointer(unexportedField.UnsafeAddr())).Elem()
	return exported.Interface()
}


func main(){
	ctx := context.WithValue(
		context.WithValue(
			context.WithValue(context.Background(), "make", "kawaksaki"),
			"model",
			"klx",
		),
		"cc",
		250,
	)

	keys, values := GetAllValues(ctx)
	fmt.Printf("keys: %s\n", keys)
	fmt.Printf("values: %s\n", values)
}
