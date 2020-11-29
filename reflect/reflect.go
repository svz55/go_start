package main

import (
	"fmt"
	"reflect"
)

// func String() string
// fmt.Stringer

type S struct{}

func (s S) String() string {
	return ""
}

func main() {
	var y fmt.Stringer

	fmt.Println(IsStringer(42))
	fmt.Println(IsStringer("123"))

	fmt.Println(IsStringer(S{}))

	fmt.Println(IsStringerReflect(42, y))
	fmt.Println(IsStringerReflect("123", y))

	fmt.Println(IsStringerReflect(S{}, y))
}

func IsStringer(x interface{}) bool {
	_, ok := x.(fmt.Stringer)
	return ok
}

func IsStringerReflect(x interface{}, y interface{}) bool {
	vx := reflect.TypeOf(&x)
	vy := reflect.TypeOf(&y)
	return vx.Kind() == vy.Kind()
}
