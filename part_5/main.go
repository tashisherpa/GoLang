package main

import (
	"fmt"
)
/*
	Struct:
		Structs are a way of defining your own type. Creating a struct:
			
			type <name of our type> struct {}
			
		Structs can hold mixed types in the form of fields which we can define by name 
			
			type gasEngine struct{
				milePerGallon uint8
				gallons uint8
			}	
*/

type gasEngine struct{
	milePerGallon uint8
	gallons uint8
	ownerInfo owner
}

/*

	The fields can be anything you want even another struct Example:
		type gasEngine struct{
			milePerGallon uint8
			gallons uint8
			ownerInfo owner
		}
			
		or 

		type gasEngine struct{
			milePerGallon uint8
			gallons uint8
			owner
		}
		doing the above where we put the owner type directly instead of having
		ownerInfo field with a name subfield. We are adding the subfields directly.
		In other words we can just use the myEngine.name syntax. This will look like:
			myEngine{
				milePerGallon: 25
				gallons: 15
				name: "Alex"
			}

		type owner struct{
			name string
		}

		We can also do this with any types, for example:
		
		type gasEngine struct{
			milePerGallon uint8
			gallons uint8
			owner
			int
		}
		now we have a field with int type which also of type int

		var myEngine gasEngine = gasEngine{25,15,owner{"Alex"}, 10}
		myEngine{
			milePergallon: 25
			gallons: 15
			name: "Alex"
			int: 10
		}
	
*/

type owner struct{
	name string
}
	
/*
	Struct also have concept of methods which we can use as well.
	These are functions which are directly tied to the struct and have
	access to the struct instance itself. 
	Example: We want to create a function that calculates the miles left
	in the gas tank

	func (e gasEngine) miles() uint8{
		return e.gallons * e.milePerGallon
	}

	expect for the "(e gasEngine)" part a method is just like a function
	in this case we return a uint8. What we are doing with (e gasEngine)
	is that we are assigning this function to the gasEngine type. Now this
	function has access to the fields and even other methods that we have
	assigned to the gasEngine type

	Note: this is similar to classes where we are instantiating a class and
	calling one of its methods
*/

func (e gasEngine) milesLeft() uint8{
	return e.gallons*e.milePerGallon
}

/*
	Now we can pass in the new type to functons like this:

	func canMakeIt(e gasEngine, miles uint8){
		if miles <= e.milesleft(){
			fmt.Println("You can make it there!")
		}else{
			fmt.Println("Need to fuel up first!")
		}
	}

	this function takes in a gasEngine type and a miles parameter
	and check if we can drive that distance
*/

/*
	Suppose now we also have another type of engine
	type electricEngine struct{
		milePerKiloWattHour uint8
		kiloWattHour uint8
	}
	
	and has similar milesLeft function
	
	func (e electricEngine) milesLeft() uint8{
	 	return e.milePerKiloWattHour * e.kiloWattHour
	}
*/

type electricEngine struct{
	milePerKiloWattHour uint8
	kiloWattHour uint8
}

func (e electricEngine) milesLeft() uint8{
	return e.kiloWattHour * e.milePerKiloWattHour
}

//Currently our canMakeIt function only takes in gasEngine type
// func canMakeIt (e gasEngine, miles uint8){
// 	if miles <= e.milesLeft(){
// 		fmt.Println("You can make it there!")
// 	}else{
// 		fmt.Println("Need to fuel up first!")
// 	}
// }
//But what if we want this more General and allow it to take any type of engine
//type

/*
	This is where the interfaces come in. lets define an interface and how they can help us with
	above issue. We use similar syntax to defining a struct but with interface keyword instead of struct

	type <name of the type> interface{}

	also in our canMakeIt() function, we really need is the mileLeft() method which takes no parameters
	and returns an uint8. This is called the method signature. We can specify this signature within the inferface like so:

	type engine interface{
		milesLeft() uint8
	}
*/

type engine interface{
	milesLeft() uint8
}

//Instead of our method here requiring the e parameter to be gasEngine we replace it with our
//interface like so. This function now can take in anything for this parameter with only requirement
// the object has a milesLeft() method with the signature we specified in our interface. 
// This way we can apply this function to a wider range of engine types 
func canMakeIt (e engine, miles uint8){
	if miles <= e.milesLeft(){
		fmt.Println("You can make it there!")
	}else{
		fmt.Println("Need to fuel up first!")
	}
}


func main(){
	fmt.Println("Struct and Interface")
	/*
	because we havent defined the miles per gallon or gallons
	field yet this is a zero valued struct, meaning default values are
	set for the fields
	
	myEngine{
		milePerGallon: 0
		gallons: 0
	}

	One way to assign value to the struct is to use the struct literal syntax like
	var myEngine gasEngine = gasEngine{milePerGallon:25, gallons:15}
	*/

	// assigning using struct literal
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}
	/*
	We can also omit the field names and the fields are asigned values
	in order 

	var myEngine gasEngine = gasEngine{25,15}

	We can also assign value by name directly like:
	myEngine.gallons = 20

	*/
	fmt.Println(myEngine,myEngine.milePerGallon, myEngine.gallons)
	
	/*
		Because we now have another struct as a field in our gasEngine Struct,
		we create an owner info field like this when assigning

		var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}
	*/ 

	//ownerInfo can be accessed like below
	fmt.Println(myEngine.ownerInfo.name)

	/*
	Note: you can also declare anonymous structs where you aren't a name type
	like above for gasEngine we did. With an anonymous struct we have to define
	and initialize it in the same location. Example:

	var myEngine = struct{
		milePerGallon uint8
		gallons uint8
		}{25,15}

	Main difference is this is not reusable if we want to create another struct
	like this we will have to rewrite the definition
	
	var myEngine2 = struct{
		milePerGallon uint8
		gallons uint8
		}{40,15}
*/
	var animal = struct{
		name string
		sound string
	}{"dog", "bark"}

	var animal1 = struct{
		name string
		sound string
	}{"cow", "moo"}

	fmt.Printf("%v's %v\n", animal.name, animal.sound)
	fmt.Printf("%v's %v\n", animal1.name, animal1.sound)

	//calling the milesLeft() function in our main
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	var myGasEngine gasEngine = gasEngine{25,15,owner{"Tashi"}}
	canMakeIt(myGasEngine, 50)

	var myElectricEngine electricEngine = electricEngine{25, 15}
	canMakeIt(myElectricEngine, 50)

}