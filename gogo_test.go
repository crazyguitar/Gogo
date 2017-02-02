package gogo

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sort"
	"time"
)

func FuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

const README = "README.md"

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

	var fileName string = README

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

// A Timer
func StartTimer() func() {
	// get function names
	f := FuncName()

	t := time.Now()
	log.Println(f, "started")
	return func() {
		d := time.Now().Sub(t)
		log.Println(f, "took", d)
	}
}

// An example about timeing
func ExampleTimer() {
	stop := StartTimer()
	defer stop()

	time.Sleep(1 * time.Second)
	// Output:
}

func PrintIntSlice(a ...int) {
	str := fmt.Sprintln(a)
	fmt.Print(str)
}

func PrintArbitraryArgs(v ...interface{}) {
	str := fmt.Sprintln(v...)
	fmt.Print(str)
}

// An example about arbitrary number of parameters
//
// ref: https://golang.org/doc/effective_go.html#printing
//
func ExampleArbitraryNumArgs() {
	s := []int{9487, 9527, 5566}

	PrintIntSlice(s...)
	PrintArbitraryArgs("Hello ", "Go!", s)
	// Output:
	// [9487 9527 5566]
	// Hello  Go! [9487 9527 5566]
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

// An example about using ioutil read a file
//
// ref: http://stackoverflow.com/q/1821811
//
func ExampleReadFile() {

	// check README file exists
	_, err := os.Stat(README)
	if err == nil {
	} else if os.IsNotExist(err) {
		fmt.Printf("Please Create a '%s'\n", README)
		panic(err)
	} else {
		panic(err)
	}

	// read README file content
	c, err := ioutil.ReadFile(README)
	if err != nil {
		panic(err)
	}
	fmt.Print(len(c) != 0)
	// Output:
	// true
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

type Geek struct {
	name   string
	skills []string
}

type Hacker struct {
	Geek
	arrest bool
}

func (g *Geek) LearnSkill(skill string) {
	geek := *g
	geek.skills = append(geek.skills, skill)
	*g = geek
}

func (h *Hacker) LearnSkill(skill string) {
	hacker := *h
	if hacker.arrest {
		return
	}
	hacker.skills = append(hacker.skills, skill)
	*h = hacker
}

// An example about method
func ExampleGoMethod() {
	geek := Geek{name: "GoGeek", skills: []string{"C", "C++"}}
	hacker := Hacker{
		Geek:   Geek{name: "GoHacker", skills: []string{"C", "Go"}},
		arrest: false}

	// Geek learn the skills
	geek.LearnSkill("Python")

	// hacker learn the skills
	hacker.LearnSkill("Javascripts")
	hacker.arrest = true // under arrest
	hacker.LearnSkill("Python")

	// show the skills
	for _, skill := range geek.skills {
		fmt.Printf("Geek have skill: %s\n", skill)
	}
	for _, skill := range hacker.skills {
		fmt.Printf("Hacker have skill: %s\n", skill)
	}
	// Output:
	// Geek have skill: C
	// Geek have skill: C++
	// Geek have skill: Python
	// Hacker have skill: C
	// Hacker have skill: Go
	// Hacker have skill: Javascripts
}

func Fib(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	c <- a
}

// An example about goroutine
func ExampleGoroutine() {
	var out []int

	// create a channel
	c := make(chan int)

	// calculate Fib numbers
	in := []int{30, 20, 35, 10, 50}
	for _, i := range in {
		go Fib(i, c)
	}
	// retrive Fib results
	for i := 0; i < len(in); i++ {
		out = append(out, <-c)
	}

	// print results
	sort.Ints(out)
	for _, v := range out {
		fmt.Println(v)
	}
	// Output:
	// 55
	// 6765
	// 832040
	// 9227465
	// 12586269025
}

type DuckProto struct {
	age     int
	friends []string
}

type DonaldDuck struct {
	DuckProto
	family []string
}

type DaffyDuck struct {
	DuckProto
	species string
}

type Duck interface {
	walk()
	swim()
	talk()
}

func (d DonaldDuck) walk() {
	fmt.Println("Donald Duck walking...")
}

func (d DonaldDuck) swim() {
	fmt.Println("Donald Duck swimming...")
}

func (d DonaldDuck) talk() {
	fmt.Println("I'm Donald Duck.")
	fmt.Println("I'm ", d.age)
	fmt.Println("My friends: ", d.friends)
	fmt.Println("My family: ", d.family)
}

func (d DaffyDuck) walk() {
	fmt.Println("Daffy Duck walking...")
}

func (d DaffyDuck) swim() {
	fmt.Println("Daffy Duck swimming...")
}

func (d DaffyDuck) talk() {
	fmt.Println("I'm Daffy Duck.")
	fmt.Println("I'm ", d.age)
	fmt.Println("My species: ", d.species)
	fmt.Println("My friends: ", d.friends)
}

// An example about duck typing
func ExampleDuckType() {
	var duck Duck

	donald := DonaldDuck{
		DuckProto{age: 83, friends: []string{"Mickey", "Minnie", "Goofy"}},
		[]string{"McDuck", "Huey", "Dewey", "Louie"}}

	daffy := DaffyDuck{
		DuckProto{age: 80, friends: []string{"Bunny", "Porky"}},
		"American black duck"}

	// donald satisifid duck type
	duck = donald
	duck.walk()
	duck.swim()
	duck.talk()

	fmt.Println("---")

	// daffy satisifid duck type
	duck = daffy
	duck.walk()
	duck.swim()
	duck.talk()

	// Output:
	// Donald Duck walking...
	// Donald Duck swimming...
	// I'm Donald Duck.
	// I'm  83
	// My friends:  [Mickey Minnie Goofy]
	// My family:  [McDuck Huey Dewey Louie]
	// ---
	// Daffy Duck walking...
	// Daffy Duck swimming...
	// I'm Daffy Duck.
	// I'm  80
	// My species:  American black duck
	// My friends:  [Bunny Porky]
}
