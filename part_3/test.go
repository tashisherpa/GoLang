package main

import (
	"fmt"
	"time"
)

func test() {
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)
	//Ran the test three times
	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n)) //14.593, 15.5894, 16.0205ms
	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice2, n)) //3.4853, 4.0613, 3.1845ms
}

func timeLoop(slice []int, n int) time.Duration{
	var t0 = time.Now()
	for len(slice)<n{
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
