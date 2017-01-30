package gogo

import (
	"fmt"
	"os"
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

// An example about "defer"
func ExampleDefer() {

	var fileName string = "README.md"

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	// finally, we need to close file
	defer func() {
		fmt.Printf("Close file\n")
		err := file.Close()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Close done\n")
	}()

	fmt.Printf("Open file: '%s' success\n", fileName)
	// Output:
	// Open file: 'README.md' success
	// Close file
	// Close done
}

// An example about function in Go
func ExampleFuncPtr() {
	// declare a function pointer
	var fPtr func(string) string

	// assign a function to fPtr
	fPtr = func(str string) string {
		return str
	}

	ret := fPtr("Hello Go!")
	fmt.Println(ret)
	// Output:
	// Hello Go!
}

// An example about function collection
func ExampleFuncCollection() {

	// Using map
	m := make(map[string]func())
	m["func1"] = func() {
		fmt.Println("Run func1")
	}
	m["func2"] = func() {
		fmt.Println("Run func2")
	}
	m["func1"]()
	m["func2"]()

	// using slice
	var s []func()
	s = append(s, func() {
		fmt.Println("Run Foo")
	})
	s = append(s, func() {
		fmt.Println("Run Bar")
	})
	for _, f := range s {
		f()
	}
	// Output:
	// Run func1
	// Run func2
	// Run Foo
	// Run Bar
}

func Done(str string) {
	fmt.Printf("'%s' Done\n", str)
}

// An example about callback function
func ExampleCallback() {
	// declare a function ptr
	var fPtr func(func(string))

	fPtr = func(callback func(string)) {
		callback("fPtr")
	}

	fPtr(Done)
	// Output:
	// 'fPtr' Done
}

type person struct {
	name string
	age  int
}

type coder struct {
	person
	skills []string
}

// An example about struct
func ExampleStruct() {

	// new a coder ptr
	geek := new(coder)
	geek.name = "golang"
	geek.age = 10
	geek.skills = append(geek.skills, "Go")
	geek.skills = append(geek.skills, "C")
	geek.skills = append(geek.skills, "Python")

	fmt.Printf("I'm \"%s\"\n", geek.name)
	for _, skill := range geek.skills {
		fmt.Printf("---> I have skill: '%s'\n", skill)
	}

	// anonymous struct
	var hacker struct {
		info   coder
		arrest bool
	}
	hacker.info.name = "anonymous"
	hacker.info.age = -1
	hacker.info.skills = append(hacker.info.skills, "unknown")
	hacker.arrest = false

	fmt.Printf("Hacker: \"%s\"\n", hacker.info.name)
	fmt.Println("Arrest: ", hacker.arrest)

	// anonymous struct template data
	data := struct {
		company string
		title   string
	}{
		"anonymous group",
		"Hacker & Geek",
	}
	fmt.Printf("Work company: \"%s\"\n", data.company)
	fmt.Printf("Job title: \"%s\"\n", data.title)

	// Output:
	// I'm "golang"
	// ---> I have skill: 'Go'
	// ---> I have skill: 'C'
	// ---> I have skill: 'Python'
	// Hacker: "anonymous"
	// Arrest:  false
	// Work company: "anonymous group"
	// Job title: "Hacker & Geek"
}
