package main

import (
	"os"
	"runtime/pprof"
)

type NumbersInt64 struct {
	numbers []int64
}

func (n *NumbersInt64) initialize() {
	for i:=0 ; i<=1000000; i++ {
		n.numbers = append(n.numbers, int64(i*2))
	}
}

func (n *NumbersInt64) addNums() {
	slice := n.numbers
	for idx, n := range slice {
		slice[idx] = n + 10
	}
}

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	numbers := &NumbersInt64{}
	numbers.initialize()
	numbers.addNums()
}


