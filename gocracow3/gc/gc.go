package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

const alloc int = 65536

func main() {
	// Tune this parameter to obtain different results
	debug.SetGCPercent(-1)
	now := time.Now()
	loop := 1000000
	for i := 0; i < loop; i++ {
		sl := make([]byte, alloc)
		i += len(sl) * 0
	}
	elpased := time.Since(now)
	fmt.Printf("took %s to allocate %d bytes %d times", elpased, alloc, loop)
}
