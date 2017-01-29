package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type Func func()

func GetFunName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// An example about array operations
func Arr() {
	// Array
	var arr = [5]int{183, 9572, 5566, 9487, 7788}
	fmt.Printf("arr = %v\n", arr)
	fmt.Printf("arr[:] = %v\n", arr[:])
	fmt.Printf("arr[1:3] = %v\n", arr[1:2])
	fmt.Printf("arr[2:] = %v\n", arr[2:])
	fmt.Printf("arr[:3] = %v\n", arr[:3])
}

// An example about slice operations
func Slice() {
	// declare a slice
	var s []int
	s = make([]int, 3, 3)
	s[0] = 9527
	s[1] = 5566
	s[2] = 9487

	for i, v := range s {
		fmt.Printf("s[%d] = %d\n", i, v)
	}

	// or
	ss := []int{9527, 5566, 9487}
	for i, v := range ss {
		fmt.Printf("ss[%d] = %d\n", i, v)
	}
}

// An example about map operations
func Map() {

	// declare a map
	var m map[string]string

	// create a map
	m = make(map[string]string)

	m["FOO"] = "foo"
	m["BAR"] = "bar"
	m["BAZ"] = "baz"
	for k, v := range m {
		fmt.Printf("m[%s] = %s\n", k, v)
	}
}

// Main function
func main() {
	var funcMap map[string]Func

	funcMap = make(map[string]Func)

	// add functions to map
	funcMap[GetFunName(Arr)] = Arr
	funcMap[GetFunName(Map)] = Map
	funcMap[GetFunName(Slice)] = Slice

	for k, v := range funcMap {
		fmt.Printf("---> Example: %s\n\n", k)
		v()
		fmt.Printf("\n")
	}
}
