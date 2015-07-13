package main

import (
	"fmt"
	"reflect"
)

type someClass struct {
	Initialized0 bool
	Initialized1 bool
}

func (t *someClass) Init() (bool) {
	fmt.Printf("*someClass.Init()\n");
	t.Initialized0 = true

	return true
}

func (t someClass) StrangeInit() (bool) {
	fmt.Printf(" someClass.Init()\n");
	t.Initialized1 = true

	return true
}

func callMethodByName(obj interface{}, methodName string) (bool, error) {
	objValue := reflect.ValueOf(obj)

	methodValue := objValue.MethodByName(methodName)
	if (!methodValue.IsValid()) {
		return false, fmt.Errorf("There's no such method \"%v\" in obj of type \"%v\"", methodName, objValue.Type().Name())
	}
	result := methodValue.Call([]reflect.Value{})

	return result[0].Interface().(bool), nil
}

func main() {
	var obj someClass

	// Practicing in "MethodByName" (reflect):

	res,err := callMethodByName(&obj, "Init")	// You cannot call "StrangeInit" on "&obj" and "Init" on "obj".
	if (err != nil) {
		fmt.Printf("Cannot call the method \"Init\": %v\n", err.Error());
		return;
	}
	fmt.Printf("res0: %v\n", res);

	res,err  = callMethodByName( obj, "StrangeInit")
	if (err != nil) {
		fmt.Printf("Cannot call the method \"Init\": %v\n", err.Error());
		return;
	}
	fmt.Printf("res1: %v\n", res);

	fmt.Printf("\nobj == %v\n", obj)
}
