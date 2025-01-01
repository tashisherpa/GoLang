package main

import (
	"fmt"
	"time"
)

func withoutGoRoutine(){
	t1 := time.Now()
	for i:=0; i <1000; i++{
		dataCall()
	}
	fmt.Printf("\nTotal execution time with delay of 2sec without goroutine: %v", time.Since(t1))
}

func dataCall(){
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
}