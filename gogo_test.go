package gogo

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
)

type Func func()

func GetFunName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// An example about array operations
func ExampleArr() {
	// Array
	var arr = [5]int{183, 9572, 5566, 9487, 7788}
	fmt.Printf("arr = %v\n", arr)
	fmt.Printf("arr[:] = %v\n", arr[:])
	fmt.Printf("arr[1:3] = %v\n", arr[1:2])
	fmt.Printf("arr[2:] = %v\n", arr[2:])
	fmt.Printf("arr[:3] = %v\n", arr[:3])
	// Output:
	// arr = [183 9572 5566 9487 7788]
	// arr[:] = [183 9572 5566 9487 7788]
	// arr[1:3] = [9572]
	// arr[2:] = [5566 9487 7788]
	// arr[:3] = [183 9572 5566]
}

// An example about slice operations
func ExampleSlice() {
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
	// Output:
	// s[0] = 9527
	// s[1] = 5566
	// s[2] = 9487
	// ss[0] = 9527
	// ss[1] = 5566
	// ss[2] = 9487

}

// An example about map operations
func ExampleMap() {

	// declare a map
	var m map[string]string

	// create a map
	m = make(map[string]string)

	m["FOO"] = "foo"
	m["BAR"] = "bar"
	m["BAZ"] = "baz"
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	// sort strings
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("m[%s] = %s\n", k, m[k])
	}
	// Output:
	// m[BAR] = bar
	// m[BAZ] = baz
	// m[FOO] = foo
}
