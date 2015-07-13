package main

import (
	"fmt"
	"reflect"
)

type someClass struct {
	Initialized0 bool
	Initialized1 bool
}

func (t *someClass) Init() {
	fmt.Printf("*someClass.Init()\n");
	t.Initialized0 = true
}

func (t someClass) StrangeInit() {
	fmt.Printf(" someClass.Init()\n");
	t.Initialized1 = true
}

func callMethodByName(obj interface{}, methodName string) (error) {
	objValue := reflect.ValueOf(obj)

	methodValue := objValue.MethodByName(methodName)
	if (!methodValue.IsValid()) {
		return fmt.Errorf("There's no such method \"%v\" in obj of type \"%v\"", methodName, objValue.Type().Name())
	}
	methodValue.Call([]reflect.Value{})

	return nil
}

func main() {
	var obj someClass

	// Practicing in "MethodByName" (reflect):

	err := callMethodByName(&obj, "Init")	// You cannot call "StrangeInit" on "&obj" and "Init" on "obj".
	if (err != nil) {
		fmt.Printf("Cannot call the method \"Init\": %v\n", err.Error());
		return;
	}

	err  = callMethodByName( obj, "StrangeInit")
	if (err != nil) {
		fmt.Printf("Cannot call the method \"Init\": %v\n", err.Error());
		return;
	}

	fmt.Printf("%v\n", obj)
}
