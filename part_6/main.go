package main

import "fmt"

func main() {
	fmt.Println("Pointer")
	/*
	What are Pointers and how are they used in GoLang?

	- pointer variable
	var p = 0x140000b2020

	- Normal variable
	var i = 123
	var f = 2.3

	Pointers are special type these are variables which store
	memory locations rather than values like intergers or floats
	
	To create a pointer we create a "*" symbol 
	*/

	var p *int32  = new(int32)
	// this variable p will hold the memory address of an int32 value

	//a regular int32 variable
	var i int32
	fmt.Println(p)
	fmt.Println(i)

	/*
	How these variables might look like in momory

	varible		Memory Location
	p 	nil			0x1b00
					0x1b01
					0x1b02
					0x1b03
					0x1b04
					0x1b05
					0x1b06
					0x1b07
	i	0			0x1b08
					0x1b09
					0x1b0a
					0x1b0b
					0x1b0c
					0x1b0d
					0x1b0e
					0x1b0f

	p is going to store a pointer or memory address which itself
	takes up to 32 or 64 bits depending on your OS. In above table
	its 64 bits or 8 byes.i
	The p pointer is initially nil because we havent initailized it
	yet in other words this pointer doesn't yet point to anything or
	another way to this is this pointer does not have an address
	assigned to it in which to store an int32 value yet.
	To give this pointer an address we can use the built-in new()
	function. What this does is, it gives us back a free memory
	location which is 32 bits wide which P can use to store an int32
	value
	example: var p *int32 = new(int32)
	
	varible		Memory Location
	p 	0x1b0c		0x1b00
					0x1b01
					0x1b02
					0x1b03
					0x1b04
					0x1b05
					0x1b06
					0x1b07
	i	0			0x1b08
					0x1b09
					0x1b0a
					0x1b0b
		0			0x1b0c
					0x1b0d
					0x1b0e
					0x1b0f

	now you can see that p stores a memory location 
	i.e: it points to the "0x1b0c" memory location.

	Note: Pointers are special type of variables, this is true
	but they still share a lot in common with regular variables.
	They still have memory address themselves and they store a value
	at that address. In the case of a pointer this is another memory
	address.

	If we want to get the value stored at this memory location
	we can use the star symbol like shown below:
		fmt.Printf("The value of p points to is: %v", *p)
		output: The value of p points to is 0
	This is called dereferencing the pointer. We get back 0 because
	this is the default value of int32. In fact, when you initialize a
	pointer with a memory location, it zeros out the memory location.
	In other words, it sets the value at the memory location to the
	zero value of that type. For example: 0 for int32, "" for string
	and false for boolean.

	To change the value stored at the memory location of a pointer, we
	use  star notation and assign it a value. Example:
		*p = 10 
	the line above, says to set the value at the memory location P is
	pointing to 10 

	Note: the star notation "*" does double duty which may be a little
	confusing. 
	var p *int32 = new(int32) here we use to tell the compiler that we
	want to initialize a pointer

	but in the following codes, we tell the compiler that we want to
	reference the value of the pointer.

	fmt.Printf("The value of p points to is: %v", *p)
	*p = 10

	These are two separate roles the star syntax "*" has that you should
	keep in mind. 
	
	Another common source of headaches is trying to get or set the value
	of a nil pointer so if you don't call the new() function. We don't see
	any compiler errors but when we run the code we get a nil pointer
	exception another word is a runtime error
		var point *int32
	What is happening here is we did not assign a memory address to our
	pointer so we obviously can't get the value at a memory address that doesnt
	exist so to make sure our pointer isn't nil before trying to assign a values
	to it.

	Next, we can also create a pointer from the address of another variable using
	"&" ampersand symbol like:
	p = &i
	The "&" symbol here means that we want the memory address of the varible
	not its value so now P refers to  memory address of i. In other words,
	p and i reference the same int32 value in memory.

	Note: if we use the star notation to change the value of p, the value of i
	is also changed
		p = &i
		*p = 1
	varible			Memory Location
		p 	0x1b08		0x1b00
						0x1b01
						0x1b02
						0x1b03
						0x1b04
						0x1b05
						0x1b06
						0x1b07
		i	1			0x1b08
						0x1b09
						0x1b0a
						0x1b0b
						0x1b0c
						0x1b0d
						0x1b0e
						0x1b0f
	This is different from when using a regular variable for example
	var k int32 = 2
	i = k

	in this code above what our program will do is copy the value of K into
	i's memory location
	varible			Memory Location
		p 	0x1b08		0x1b00
						0x1b01
						0x1b02
						0x1b03
						0x1b04
						0x1b05
						0x1b06
						0x1b07
		i	2			0x1b08
						0x1b09
						0x1b0a
						0x1b0b
		k	2			0x1b0c
						0x1b0d
						0x1b0e
						0x1b0f
	*/
	fmt.Printf("The value of p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v\n", i)
	p = &i
	*p = 1
	fmt.Printf("The value of p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v\n", i)

	var k int32 = 2
	i = k
	k = 3
	fmt.Printf("The value fo p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v\n", i)
	fmt.Printf("The value of k is: %v\n", k)

	/*
		The main exception of this copy behavior of non-pointer varuables
		is when working with slices
		Lets say we copy a slice in the regular way without using pointer

		var slice = []int32{1,2,3}
		var sliceCopy = slice

		lets now modify the sliceCopy variable and printing the slices
		sliceCopy[2] = 4
		fmt.Println(slice) >> [1,2,4]
		fmt.Println(sliceCopy) >> [1,2,4]

		we can see that actually the values of original slice has changed
		this is because under the hood slices contain pointers to an underlying
		array. Basically with slices we are just copying the pointers when we
		do this. So, both varaibles of our slices refer to same data. 
	*/

	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	/*
	Using pointers in functions. These two go really nicely together.
	Lets see why?

	example:

	this is a fucntion that takes in float64 array of size 5 and squares all the values
	func square(thing2 [5]float64) [5]float64{
		for i:= range thing2{
			thing2[i] = thing2[i]*thing2[i]
		}
		return thing2
	}

	and print the memory addresses of the arrays. We can see that there are two different
	memory address meaning these are two different arrays therefore we can modify the values
	of thing2 without affecting the values of thing1 in our main function. But we are also 
	doubling our memory usuage of the variables passed in because we are creating
	copies for use in our function. 
	
	so, we are potentially using way more memory than we need
	well not anymore. Instead let's use pointers. Our function taking a pointer to an array
	instead

	func square(thing2 *[5]float64) [5]float64{
		fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)
		for i := range thing2{
			thing2[i] = thing2[i]*thing2[i]
		}
		return thing2
	}

	after we can see that both the memory address for thing1 and thing2 are the same.

	Note: Pointer are really useful in large parameters so you don't have to create 
	copies of the data everytime we call a function, wasting time and memory. Instead 
	you can pass in the pointer to the data. The only thing to be mindful of now is
	since thing1 and thing2 refer to the same array changing the values of thing2 means
	the values of thing1 also changes.
	*/

	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1) //0xc0001280c0
	//without the pointer of array in parameter
	//var result [5]float64 = square(thing1)

	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
}

//without pointer
// func square(thing2 [5]float64) [5]float64{
// 	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2) //0xc0001280f0
// 	for i := range thing2{
// 		thing2[i] = thing2[i]*thing2[i]
// 	}
// 	return thing2
// }


//with pointer
func square(thing2 *[5]float64) [5]float64{
	fmt.Printf("\nThe memory location of the thing2 array is: %p", thing2) //0xc00000c3f0
	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return *thing2
}
