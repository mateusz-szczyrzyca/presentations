package pointer_vs_value

import (
	"testing"
)

type ComplicatedStruct struct {
	S1, S2, S3, S4, S5 string
	A1, A2, A3, A4, A5 int64
	F1, F2, F3, F4, F5 float64
}

func BenchmarkPassingValue(b *testing.B) {
	v := ComplicatedStruct{
		"a","b", "c", "d", "e",
		1,2,3,4,5,
		1.1, 1.2, 1.3, 1.4, 1.5,
	}

	f := func(s ComplicatedStruct) {
		s.A1 = 10
		s.A2 = 20
		s.A3 = 30
		s.A4 = 40
		s.A5 = 50
	}

	for i := 0; i < b.N*10; i++ {
		f(v)
	}
}

func BenchmarkPassingPointer(b *testing.B) {
	v := ComplicatedStruct{
		"a","b", "c", "d", "e",
		1,2,3,4,5,
		1.1, 1.2, 1.3, 1.4, 1.5,
	}

	f := func(s *ComplicatedStruct) {
		s.A1 = 10
		s.A2 = 20
		s.A3 = 30
		s.A4 = 40
		s.A5 = 50
	}

	for i := 0; i < b.N*10; i++ {
		f(&v)
	}
}
