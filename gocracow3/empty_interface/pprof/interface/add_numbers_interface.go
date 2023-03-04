package main

import (
	"os"
	"runtime/pprof"
)

type NumbersEmptyInterface struct {
	numbers []interface{}
}

func (n *NumbersEmptyInterface) initialize() {
	for i:=0 ; i<=1000000; i++ {
		n.numbers = append(n.numbers, i*2)
	}
}

func (n *NumbersEmptyInterface) addNums() {
	slice := n.numbers
	for idx, n := range slice {
		slice[idx] = n.(int) + 10
	}
}

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	numbers := &NumbersEmptyInterface{}
	numbers.initialize()
	numbers.addNums()
}
