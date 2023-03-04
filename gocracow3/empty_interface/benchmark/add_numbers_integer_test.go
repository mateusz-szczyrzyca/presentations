package benchmark

import "testing"

type NumbersInt64 struct {
	numbers []int64
}

func (n *NumbersInt64) initialize() {
	for i:=0 ; i<=1000; i++ {
		n.numbers = append(n.numbers, int64(i*2))
	}
}

func (n *NumbersInt64) addNums() {
	slice := n.numbers
	for idx, n := range slice {
		slice[idx] = n + 10
	}
}

type NumbersEmptyInterface struct {
	numbers []interface{}
}

func (n *NumbersEmptyInterface) initialize() {
	for i:=0 ; i<=1000; i++ {
		n.numbers = append(n.numbers, i*2)
	}
}

func (n *NumbersEmptyInterface) addNums() {
	slice := n.numbers
	for idx, n := range slice {
		slice[idx] = n.(int) + 10
	}
}

func BenchmarkAddIntegerNumbers(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		numbers := &NumbersInt64{}
		numbers.initialize()
		numbers.addNums()
	}
}

func BenchmarkAddEmptyInterfaceNumbers(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		numbers := &NumbersEmptyInterface{}
		numbers.initialize()
		numbers.addNums()
	}
}

