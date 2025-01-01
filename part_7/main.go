package main

import (
	"fmt"
	"time"
	"sync"
)

/*
Goroutines:
	they are a way to launch multiple functions and have them execute
	concurrently.

	Note: Concurrency != parallel execution

	concurrency means that I have multiple tasks in progress at the same time.
	One way we can do this is by jumping back and forth from working on one
	task to another.
	Example: let's say task one involves the database call which takes
	3 seconds to return a data. In concurrent programing while I'm waiting for
	the database to response the CPU can move on to working on task two in the
	meantime. And when I get a response maybe we move back to finish up task1
	and move back to task2 and finish it up.

	Another way to achieve concurrencies through parallel execution instead of
	having 1 CPU core working on these two tasks we can have two CPU cores
	then, their execution can happen simultaneously. The execution here is still
	concurrent because we have multiple tasks in progress at the same time but
	also these tasks are running in parallel.

	Note: We can see that a program maybe running concurrently but may not necessarily
	be executing tasks in parallel.

	In practice, through go we usually do achieve some level of parallel execution
	using goroutines as long as you have multi-core CPU which we probably do.

*/

//creating a wait group
var wg = sync.WaitGroup{}

//creating a mutex
var m = sync.RWMutex{}

//Mock Database returns
var dbData = []string{"id1","id2","id3","id4","id5"}

//creating a slice to store results from the database
var results = []string{}

func main() {
	// t0 := time.Now()
	//this is calling the database sequentially one by one which will take
	// an avarage of 5 seconds to complete
	// for i := 0; i < len(dbData); i++{
	// 	dbCall(i)
	// }
	// fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	
	/*
	A better way to do this is to let these database calls run concurrently. 
	To do this we use the go keyword in front of the function we want to run
	concurrently. Now our program wont wait for this function to complete
	rather it'll keep moving on to the next step in the loop
	*/
	t1 := time.Now()
	for i := 0; i < len(dbData);i++{
		wg.Add(1) //increments the counter
		go dbCall(i)
	}
	wg.Wait() //waits for the counter to go back to 0
	fmt.Printf("\nTotal execution time: %v", time.Since(t1)) //527.8µs
	fmt.Printf("\nThe results are %v", results)
	/*
	What happened?
		Our program spawned these tasks in the background. Didn't wait for them
		to finish and then exited the program before they were complete. So we need
		a way for our program to wait until all these tasks are completed and continue
		on to the rest of the code. 
	This is where wait groups which can be import thru "sync" package. 
	What are wait groups? They are pretty much just counters whenever we spawn a go routine
	we make sure to add 1 to the counter like this:
		for i:=0; i < len(dbData);i++{
			wg.Add(1)
			go dbCall(i)
		} 
	
	and inside our function we then call the done method at the end
	func dbCall (i int){
		...
		wg.Done() //this decrements the counter
	}

	and finally we also call the wait method. This method is going to wait for the
	counter to go back to zero meaning that all the tasks have completed and the rest of
	the code will execute
	
	wg.Wait()
	*/

	/*
	Now what if instead of just printing out the results of the console we want it
	them to main function. Well first lets make a slice where we can store all the results 
	from the database.
		var result = []string{}

	In DB call function let's append the value we get back from our fake database
	and set our delay to 2 seconds. This way we can see what happens when our slices is
	modified at the same time by multiple go routines. 

	When you have multiple threads modifying the same memory location at the same
	time you can get some unexpected results.  

	For example:
		Two processes writing to the same memory location at the same time could lead
		to corrupt memory as well as a whole host of other issues.
		
	So, we really shouldn't run this code like this. Instead we can use what is called a 
	mutex to control the writing to our slice in a way that makes it safe in a concurrent
	program like ours. We can create a mutex from "sync" package like this:
		var m = sync.Mutex()
	Note: Mutex is short for Mutual exclusion
	The two main methods are the .Lock() and .Unlock() methods and we will place them
	around our code which accesses the result slice

	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()

	What happens here?
		When a goroutine reaches this lock method a check is performed to see if a lock has
		already been set by another goroutine. If it has it will wait here until the lock 
		is released and set the lock itself. Once it's done tinkering with our results array
		the lock is released again with the unlock method and now other goroutines can obtain 
		a lock as needed
	
	Note: 
		It really matters where you put the Lock() and Unlock() method.
		
		for example: this fucntion will breka the program
		func dbCall(i int){
			var delay float32 = 2000
			m.Lock()
			time.Sleep(time.Duration(delay)*time.Millisecond)
			fmt.Println("The result from the database is", dbData[i])
			results = append(results, dbData[i])
			m.Unlock()
			wg.Done()
		}
		
		this runs the code simulatiously rather than concurrent as the lock is placed before 
		Sleep method for 2 seconds. So the execution time 5 calls would be 10seconds
		
	One draw back of this sort of mutex is that it completely locks out other go routines
	to accessing a result slice now, we might want this but we might not.
	Go provides another type of mutex called a read write mutex. 
		var m = sync.RWMutex
	This has the same functionality of Mutex and Lock() and Unlovk() also works the same. But we 
	now also have a RLock() "read lock" and RUnlock() "read unlock" method as well. 
	For readability and improving code lets make another function

	Write to our result slice
	func save(result string){
		m.Lock()
		result = append(resilts, result)
		m.Unlock()
	}

	read our result slice
	func log(){
		m.RLock()
		fmt.Printf("\nThe current results are: %v", results)
		m.RUnlock()
	}

	Note: Many goroutines may hold readlocks at the same time. These read locks will
		only block code execution  m.Lock(). When a goroutine hits the line in order to
		proceed all locks must be cleared that is both full locks and read locks. This
	prevents us from accessing the slice while other goroutines are writing to or reading
	from the slice
	Summary: 
		This pattern allows multiple goroutines to read from our slice at the same time
		only blocking when write may be potentially be happening. This is contrast to what
		we saw with lock and unlock method of regular mutex. In that case, even read of the
		data can only happen one at a time.
	*/
	
	//performace test with and without go routines
	//both iterate 10000 times and have a delay of 2 seconds on each execution
	withGoRoutine() // average = ~2sec
	withoutGoRoutine() // 1000*2000 = 2000000/60 = ~33333 sec

}
/*
	Here we have a function that simulates a call to a database retrieving
	an Id from dbData list. 
*/
func dbCall(i int){
	//simulate DB call delay
	//takes a random amount of time up to 2 seconds per call
	var delay float32 = 2000
	// m.Lock() 
	time.Sleep(time.Duration(delay)*time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}
//Write to our result slice
func save(result string){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

//read our result slice
func log(){
	/*
	When a goroutine reaches the line below it checks if there's a full lock
	on the mutex. By full lock we mean the lock we called .Lock() method. If 
	this full lock exists it'll wait until it's released before continuing. 
	
	This way we are not reading while the results are potentially being written
	to. If no full lock is in place a go routine will acquire a read lock and then
	proceed with the rest of the code	
	*/
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}