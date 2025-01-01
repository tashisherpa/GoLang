package main

import (
	"fmt"
	"time"
)

func withGoRoutine(){
	t0 := time.Now()
	for i:=0; i <1000; i++{
		wg.Add(1)
		go databaseCall()
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time with delay of 2sec with goroutine: %v", time.Since(t0))
}

func databaseCall(){
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	wg.Done()
}