package main
import (
	"fmt"
)

func main(){
	test()
	fmt.Println("Arrays")
	/*
	Array:
		is a fixed length collection of data all of the same type, 
		which is indexable and stored in contiguous memory locations
		Example: 
		var intArr [3]int32
		Fixed Length:
			The [] brackets with the number indicates that the array holds
			exactly that amount of elements. Note: the lenght of the array
			cannot change once it is initialized
		
		Same Type:
			In the declaration, we specified the type to be int32, this means
			all the elements in the array are of that type. When declaring an
			default array, it is said to have default values of the element types

			i.e: fmt.println(intArr) >> [0,0,0]

		Indexable:
			we can access the ith element in array like:
				intArr[0]
			or access 1 and 2 elements like
				intArr[1:3]
			
			we can also change any value of the array by indexing as well
				intArr[1] = 123

		Contiguous in Memory:
			in our case, because int32 is 4 bytes of memory, and we have 3 elements
			Go allocates 12 bytes of continuous memory when we initialize the array.
			
			We can print out the memory location of each element using the '&' symbol like:
				fmt.Println(&intArr[0]) >> 0xc00000a100 --> first element in the first 4 bytes
				fmt.Println(&intArr[1]) >> 0xc00000a104 --> second in the next 4 bytes and so on..
				fmt.Println(&intArr[2]) >> 0xc00000a108
			
			Note: the benefit of this is that the compiler doesn't need to store memory locations for each element,
			      it just needs to know where the first byte is stored and increment by 4 to get to the next location
	*/
	var intArr [3]int32
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	//priting memory locations
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	//We can also immediately initialize the array using the following syntax
	var intArr1 [3]int32 = [3]int32{1,2,3}  
	fmt.Println(intArr1)

	// we can also use the :=
	myArr := [3]int32{1,2,3}
	fmt.Println(myArr)

	/*
	we can also ommit the three and have it be inferred by the compiler using
	the [...] syntax. This is still an array of fixed size 3 because we set three
	elements like below:
	*/
	myArr1 := [...]int32{1,2,3}
	fmt.Println(myArr1)

	fmt.Println("Slices")
	/*
	Slices:
		Related to array are slices, sclices are just wrappers around array according to
		the Go documentation. So under the hood slices are just arrays with additional
		functionality

		Example:
			var intSlice []int32 = []int32{4,5,6} 
			fmt.Println(intSlice) >> [4,5,6]
		By omitting the length value, we now have a slice. With Slices unlike arrays,
		we can add values to the slice using the built-in append function. This function
		takes in the slice as the first element and value that we want to append to the end
		as its second element. It then returns a slice with the new element appended
			intSlice = append(intSlice, 7)
			fmt.Println(intSlice) >> [4,5,6,7]
		
		So, whats actually happening here with respect to the underlying array. 
		Well, initially an array is allocated that can hold exactly three values when
		we go to append another number a check is done to see if the underlying array has 
		enough room for values. In the case above, it does not so a new array is made with
		enough capacity and the values are copied there. Example:
		[4,5,6]

		a new array = [*,*,*,7,*,*]
		after appending = [4,5,6,7,*,*] //this totally new array is returned

		
	
	*/
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	fmt.Printf("This length is %v with capacity %v", len(intSlice), cap(intSlice))
	
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("\nThis length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	/*
	The new array that is returned has the len of 4 and the capacity of 6,
	even though the capacity is 6 we are not able to access anything beyond the
	len of array as it will throw and index out of range[4] with length 4 
	*/
	//uncomment the code below to test
	//fmt.Println(intSlice[4])
	
	//We can also append multiple values to the slice by using the spread operator like:
	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)

	// another way to create a slice is to use the make function
	var intSlice3 []int32 = make([]int32, 3)
	fmt.Println(intSlice3)
	/*
	We can specity the length of the slice as well as optionally specify
	the capacity of the slice, otherwise by default the capacity will be the lenght of the slice.
	Note if you have an rough idea of the capacity your array needs its is a good idea to specify it.
	As this will help avoid your program from having to reallocate the underlying array when it needs
	store more values which can have a pretty large impact on performance.  
	*/

	fmt.Println("Map")
	/*
	Map:
		is a set of {"key":"value"} pairs, where you can look up the value using its key.
		We can declare a map using the following

		var <name> map[<type for key>]<type for value> = make(map[<type for key>]<type for value>)
	*/

	//In our case here, our keys are of type string
	// and out values are of type uint8
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	//We can also initialize a map with values immediately like
	var myMap2 = map[string]uint8{"Adam":23, "Sarah":20}
	fmt.Println(myMap2["Adam"])

	//note: if I try to get the value of a key that doesnt exist in our map
	// we get the default value of the type as our value
	fmt.Println(myMap2["Tashi"])

	/*
	We have to be careful when using map because the map will always retirn something
	even if the key does not exist.

	Luckily maps in Go also return an optional second value which is a boolean.
	This return "true" if the value is in the map else "false". We can use this as following
	*/

	var age, ok = myMap2["Tashi"]
	if ok{
		fmt.Printf("Tashi is %v years old", age)
	}else{
		fmt.Println("Invalid name")
	}

	// To delete something from map, Go has a built-in delete function where the first
	// arguement is a map and the second the key you want to delete. This will delete by reference so no return
	//values are given

	delete(myMap2, "Adam")

	//We can also iterate through maps
	/*
	Loops:
		if you want to iterate over something be a map, array or a slice, we can use a range keyword
		within our For loop like so below:
		for name:= range myMap2{
			fmt.Printf("Name: %v\n", name)
		}

		Note: when iterating over a map no order is preserved,so when we run the loop multiple times,
		we may get a different order of the keys
	*/
	for name, age:= range myMap2{
		fmt.Printf("Name: %v, Age:%v \n", name, age)
	}

	//similarly we can also iterate through arrays and slices like below
	// where i is the index and v is the value
	for i, v := range intArr{
		fmt.Printf("Index: %v, Value:%v\n",i, v)
	}

	//Go doesnt have a while loop per se but we can use the for loop achieve
	// a while loop
	var i int = 0
	for i < 10{
		fmt.Println(i)
		i = i + 1
	}

	//we can also omit the condition as well and use the break keyword inside the loop
	// for {
	// 	if i >= 10{
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	//finally the same thing can also be achieve using the following syntax
	// for i:=0; i<=10; i++ {
	// 	fmt.Println(i)
	// }
}

