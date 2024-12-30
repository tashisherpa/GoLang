package main

import (
	"errors"
	"fmt"
)
/*
Functions and control structures
 != -> not equal to
 == -> equal to 
 || -> or
 && -> and
*/

func main() {
	//define a variable and pass it to our printMe function
	var printValue string = "Hello World"
	printMe(printValue)


	var numerator = 11
	var demoninator = 2
	var result, remainder, err = intDivision(numerator, demoninator)
	/*
	we can use the printf function, we can format the strings easier using the variables.
	Example:
		fmt.Printf("The result of the interger division is %v with remainder of %v", result, remainder)
	We can add %v and it will replace the values with varaible we set at the end of the function when printing
	the first %v will be the result and the second will be remainder

	*/

	//Else if Statement
	// if err!=nil{
	// 	fmt.Println(err.Error())
	// }else if remainder == 0{
	// 	fmt.Printf("The result of the interger division is %v", result)
	// }else{
	// 	fmt.Printf("The result of the integer division %v with remainder %v", result, remainder)
	// }

	//Switch Statement: Break is implied
	switch{
	case err!=nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the interger division is %v with remainder %v", result, remainder)
	}

	//conditional switch statement
	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1,2: 
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}

}

// we can also pass in parameters
func printMe(printValue string) {
	fmt.Println(printValue)
}

/*
	this function returns and int division
	we also have to specify what type we are returning from the function like in the fucntion below.
	We can either return 1 or multiple results at the same time using the ( ) and multiple types
	
	Now we can think about what would happen if 0 is passed in as demonimator in our function. 
	If we dont have any thing to handle error we get a runtime error: interger divide by zero

	Note: A design pattern in go is that if your function can encounter errors
		 to have return type of type error along with values you are returning
	
	
*/
func intDivision(numerator int, denominator int) (int, int, error) {
	/*
	errors is another built-in type in go and if we initialize a variable of type
	error in our function, the default value is nil
	*/
	var err error
	//if the denominator is 0, return an error
	if denominator == 0 {
		/*
		To create an error type we need to import the errors package,
		where we can call the errors.New() method. This creates an error type that
		we can initialize with an error message. 

		We can't just return the error so we have to return the two ints as wells
		*/
		err = errors.New("Cannot divide by zero")
		return 0,0, err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}

/*
	handling errors in this way is a general design pattern in Go.
	If you import functions from other packages a lot of the time,
	they return an error type in addition to the other outputs.
*/