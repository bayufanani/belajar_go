package main

import (
	"fmt"
	"time"
)

func startBenchmark() {
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("tanpa prealocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("dengan prealocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 0)
	}
	return time.Since(t0)
}
