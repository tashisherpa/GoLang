package main

import (
	"fmt"
	"time"
)

/*
	All our DatabaseCall function does is sleep for 2 seconds
	we can call this 1000 times and still finish in about 2 seconds. The reason
	for this is not because we have 1000 cores to process all of these goroutines 
	in parallel but because this function isn't really doing anything after 2 seconds
	and the CPU can move on two the next goroutine.

	But if we have more computationally expensive tasks though the performance
	gain we get is going to be limited by the amount of cores we have. In these cases
	goroutines need to actually do some work and because if lets say you have 8cores in
	your machine, you can run 8 of these go routines at once and the rest of the goroutines
	have to wait until there is a CPU core available.
*/
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