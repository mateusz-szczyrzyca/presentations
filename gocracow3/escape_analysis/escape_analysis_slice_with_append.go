package main

import (
	"fmt"
)

type MyStruct struct {
	name 	string
	list 	[]string
}

func main() {
	a := MyStruct{
		"name",
		[]string{"a","b","c"},
	}

	somefunc1(a)
	if a.name == "name" {
		fmt.Printf("somefunc1 NOT MODIFIED struct\n")
	}

	somefunc2(&a)
	if a.name != "name" {
		fmt.Printf("somefunc2 MODIFIED struct\n")
	}
}

func somefunc1(a MyStruct) {
	a.name = "xyz"
}

func somefunc2(s *MyStruct){
	s.name = "xyz"
}
