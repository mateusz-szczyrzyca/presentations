package main

import (
	"sync"
	"testing"
)

type Database struct {
	FieldA    string
	FieldB    string
	FieldC    string
	Names
	Lists
}

type Names struct {
	first string
	last string
}

type Lists struct {
	ListInts []int
	ListStrings []string
}

var pool = sync.Pool{
	New: func() interface{} {
		return &Database{}
	},
}

func BenchmarkNoPool(b *testing.B) {
	var db *Database

	for n := 0; n < b.N; n++ {
		db = &Database{
			"Some text in FieldA",
			"Some text in FieldB",
			"Some text in FieldC",
			Names{
				"FirstName",
				"LastName",
			},
			Lists{
				[]int{1,2,3,4,5,6,7,8,9,10},
				[]string{"a","b","c"},
			},
		}
	}
	_ = db
}

func BenchmarkWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		db := pool.Get().(*Database)

		db.FieldA = "Changed text in FieldA"
		db.FieldB = "Changed text in FieldB"
		db.FieldC = "Changed text in FieldC"
		db.Names.first = "Name"
		db.Names.last = "Last"
		db.ListInts = []int{10,20,30,40,50}
		db.ListStrings = []string{"x", "y", "z"}

		pool.Put(db)
	}
}
