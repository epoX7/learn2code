package main

import (
	"fmt"
	"reflect"
)

func main() {
	// defining a string
	var mystring string
	// Memory Address (hex), value (no value)
	fmt.Println(&mystring, mystring)
	mystring = "example"
	// Memory Address (hex), value (example)
	fmt.Println(&mystring, mystring)
	// demonstrating a type
	fmt.Println(reflect.TypeOf(mystring))

	// defining a int
	var myint int
	// Memory Address (hex), value (no value)
	fmt.Println(&myint, myint)
	myint = 1
	// Memory Address (hex), value (example)
	fmt.Println(&myint, myint)
	// demonstrating a type
	fmt.Println(reflect.TypeOf(myint))

	// defining a bool
	var mybool bool
	// Memory Address (hex), value (no value)
	fmt.Println(&mybool, mybool)
	mybool = true
	// Memory Address (hex), value (example)
	fmt.Println(&mybool, mybool)
	// demonstrating a type
	fmt.Println(reflect.TypeOf(mybool))
}
