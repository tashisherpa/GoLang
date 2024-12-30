/*
Every go file is a part of a package and we identify the package that it belongs to at
the top of the file by typing package <name of the package>
which has to be the same for all the files within this folder.

main: this is a special package name that tells the compiler to look
	  for the entry point fucntion here.
	i.e: when creating an executable the compiler needs to know where the program should
	start from and it will look for a fucntion named main with in this main package which
	serves as the first thing that will get executed in your program

*/

/*
How to run the code:
	There are two ways to run the code:

	First way:
	Compile the code:	go build <path to the file> // this produces a binary file called main
	Run the file: ./<filename>

	Second way:
	Use this command: go run <path to the file> // does the above steps in one command
*/

package main

/*
Importing packages goes just under the package name declaration by typing import followerd by the package name
import <package name>
Also, in Go if you import a package you have to use it otherwise you will get an error
*/

import "fmt"
import "unicode/utf8"
/*
The editor will incicate that there's an error because we
havent yet created the main function which is required in the main package
This only applies to the special main package.
For Example:
	If I had another folder with a package name Blue, I wouldn't need to make
	a function named Blue
*/

/*
to create a function in Go, you use the key word func,
followed by the name. Note, this function does not tak enay parameters
*/

func main() {
	fmt.Println("Declaring a variables:")
	/*
		Declaring a variable, use the keyword var followed by the name and the type
		var <name> <type> 
		Just like imports you have to use every variable you declare otherwise throws an error. 
		Part of Go Simplicity Philosophy
		
	*/
	var intNum int = 32767 //Note: int will default to 32 or 64 bits depending on your system architecture.
	fmt.Println(intNum)

	fmt.Println("Int:")
	/*
		int
		int8
		int16
		int32
		int64
	        These are used to specify how much memory or bits you want to use to store your number.
			64 bits ints can store much larger number than 16 bit ints but take 4 times memory

		uint:
			This has all the same bit sizes as ints but only store positive integers.
			This allows you to store integers twice as large in same amount of memory

			Example: int8: (-128, 127) uint8: (0,255)

	*/

	fmt.Println("Floats:")
	var floatNum1 float32 = 12345678.9
	fmt.Println(floatNum1)

	var floatNum2 float64 = 12345678.9
	fmt.Println(floatNum2)
	/*
		We also have access to:
			float32 and float64
		similar to int these are 32 and 64 bit floating numbers used to store decimal numbers. 64 bit floats can stroe the largest and most preise decimal numbers
		but they take more memory.

		Note: In Go, there is no just float type we have specify either 32-bit or 64-bit
	*/

	fmt.Println("Operations:")
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)
    
	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)
	
	/*
		you can add, subtrack, divide and multiply two numbers and all that stuff
		but two things to note about arithmetic operations:
			1. You cannot perform operations with mixed types. Example:
				var floatNum32 float32 = 10.1
				var intNum32 int32 = 2
				var result float32 = floatNum32 + intNum32 //Error
			if we want to do these operation then we can cast one of the variable with a common type
				var result float32 = floatNum32 + (float32)intNum32 //12.1

			2. Integer division results in an integer, results are rounded down. Example:
				var intNum1 int = 3
				var intNum2 int = 2
				fmt.Println(intNum1/intNum2) >> 1
				fmt.Println(intNum1%intNum2) >> 1
			if you want to get a remainder here you can use the percent or modulo sign
	*/

	fmt.Println("String")
	var myString string = "Hello World"
	fmt.Println(myString)

	//Len returns the number of bytes in a string not the characters
	fmt.Println(len("A"))
	fmt.Println(len("γ"))

	//imported unicode/utf8 for getting the len of characters in a string
	fmt.Println(utf8.RuneCountInString("γ"))

	fmt.Println("Rune:")
	var myRune rune = 'a'
	fmt.Println(myRune) //97

	//we can also drop the var keyword and use the shorthand colon like below
	myVar := "text"
	fmt.Println(myVar)

	/*
		to declare a string type we use the keyword string
		var name string
		we can assign a value to the variable by either using a double quote or back quotes.
		var myString string = "Hello World!" >> Hello World!
		with " ", its just a single line, we can't for example continue my string onto next line but we can insert a line break
		var myString string = "Hello \nWorld!" >>
			Hello
			World
		but with ` `, we can format the strings directly
		var myString string = `Hello
		World` >>

			Hello
			World
		we can also concatenate strings together by adding strings together
		var MyString string = "Hello" + " " + "World" >> Hello World


		We can get the length of the string using the len() func
		fmt.Println(len("test")) >> 4 //Note this is not the number of characters in string but the bytes

		This is due to Go using utf-8 encoding characters outside of the vanilla ASCII character set are store with
		more than a single byte. For Example:
		fmt.Println(len("A")) >> 1
		but
		fmt.Println(len("γ")) >> 2

		so, if you expect some fancy strings in your code and you want the length of a string in the number of characters.
		Import the built in package called unicode/utf8 and called the RuneCountInString() func

		fmt.Println(utf8.RuneCountInString("γ")) >> 1
		RuneCountInString sounds like a weird name for a fucntion which finds the length of the string
		but runes are actually another data type in go and represent characters
	*/

	fmt.Println("Boolean")
	var myBoolean bool = true

	fmt.Println(myBoolean)
	//these can either be true or false

	fmt.Println("Think About what data types to use")
	/*
			For Example for int:
				the largest value that an int16 can store is 32767
				var intNum int16 = 32767
				but if you try to use any number larger than that during declaration, you will get a compiler error due to overflow error

			Also, the compiler wont show any error and run the code without any issues on following code during run time
			but will produce weird result.

			for Example for floats:
				Most floating numbers are not stored precisely in the computer so, lets say we have the following code:
				var floatNum float32 = 12345678.9
				fmt.Println(floatNum) //this will print 12345679.000000

				but if its an float64, we get the correct result back
				var floatNum float64 = 12345678.9
				fmt.Println(floatNum) //this will print 12345678.900000

			Moral of the story:
				Thionk about what data types you're using as well as what values you might encounter
				and avoid late night debugging sessions.

			Always a good idea to think about what data types to use rather than always using int64, float64 or int.intNu
			For example:
	            rgb(  , 43, 200)
				if I want to store a 256 RGB an uint8 is the best fit rather than int which could be 32 or 64 bit depending on the system.
	*/

	fmt.Println("Initializing variables")
	/*
	we can initialize a variable anywhere. By default the Go initialize them to a default value.
		The default value depends on its type
		
		Example: 
		var inputNum int16

		By defualt uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64 and rune equal to 0

		for strings its an empty string ""
		for boolean its false

		we could also create a variable but omit the type if we set the value right away
		This way the type is inferred
		Example:
		var myVar = "text"

		We can also drop the var keyword and use the shorthand colon equals
		example:
		myVar := "text"
		
		We can also initialize multiple variables at once in all the same ways we initialize the single variable
		Example:
		var var1, var2 int = 1, 2
		or 
		var1, var2 := 1, 2

		Whenever you can you should specify the types explicitly and when its not obvious
		For example:
		myVar := "Text" //this is fine as it is obvious that the value is a string

		but if 
		myVar := foo() //here we are not sure what the type of the func foo returns

		we have no idea what the type that is being returned. Adding the type when it is not obvious and is a good practice
		
		constants are atlernative to variables everything we said before also applies to constants,
		except you cant change the value once its created

		const myConst string = "Hello World"

		Also you cant just declare constants, like below
		const myConst string
		we have to initialize it with a value explicitly. Constants are useful when you dont want your code down the line to change
		the value we set. An example of usecase:
		const pi float32 = 3.1415
	*/
}
