package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	fmt.Println(myString)
	/*
		String:
			We can index the strings just like arrays using the same notation
			var indexed = myString[0]
			fmt.Println(indexed) >> 114
			When we print out the index we get a number rather than the character.
			When we print the indexed and its type we get
			fmt.Printf("%v, %T", indexed, indexed)

	*/

	//indexing the string
	var indexed = myString[0]
	fmt.Println(indexed) // 114 not the character at index

	//printing the indexed and its type: To print the type we use %T
	fmt.Printf("%v, %T\n", indexed, indexed) // 114, uint8

	/*
		now if we iterate over the string, we get something weird
		i	 v
		0	114
		1	233
		3	115
		4	117
		5	109
		6	233

		Above output we can see that it skipped 2 and second column is bunch of numbers

	*/
	for i, v := range myString {
		fmt.Println(i, v)
	}

	/*
		To understand what is going on here, we need to undertsnad utf-8 encoding
		which is how Go represents the strings on our computers.

		Remember, we have to represent strings as binary numbers in computers, one such early way
		of doing this was using ASCII encoding this uses 7 bits to encode 128 characters.
		Example: "a" = 97 in ASCII and 1100001 in binary.indexed
			ASCII (7-bits) = 129 chars

		But what do we do if we want to represent an extended set of characters like emojis
		or different countries' alphabet characters. One obvious way to solve this would be
		using more bits.
		Example:
			We can extent our character representation to use 32 bits or 4 bytes. This is exactly what
			utf-32 encoding does
				UTF-32(32-bits) = 1,114,122 chars

			but this can waste a lot of memory for many characters.
			For example:
			Characters							UTF-32							UTF-8
				a						00000000	00000000		 	   	01100001
										00000000	01100001

				😊						00000000	00000001				11110000	10011111
										11110110	00001010				10011000	10001010

				家						00000000	00000000				11100101	10101110
										01011011	10110110					10110110

			It would be nice if we did not have all the zeros there to represent "a"
			But utf-8 on the other hand seeks to solve this issue by allowing varaible
			length encoding. i.e: using the appropriate number of bytes for the character

			utf-8 uses a predefined encoding pattern which encodes information about how many bytes
			this particular character uses. Example: you can tell that a character uses one byte if it
			starts with a 0, two bytes if it starts with 110

			Example: taking this "é"
			é --> 233
			unicode characters are numbered between 128 and 2047, use two bytes and hence this pattern:
			110xxxxx 10xxxxxx

			233 in binary = 11101001, so we need to pad this number with leading 0s in order to fit
			into the utf-8 encoded representation like this:
				00011101001
			replace the x with the padded leading 0s
				11000011 10101001
			now the above is our utf-8 encoded value for é

	*/

	var resume = "résumé"
	fmt.Println(resume)
	var firstChar = resume[1]
	fmt.Printf("%v, %T\n", firstChar, firstChar) // 114, uint8
	for i, v := range resume {
		fmt.Println(i, v)
		/*
			Output of the loop:
				index 	value
				  0		114
				  1		233
				  3		115
				  4		117
				  5		109
				  6		109
		*/
	}
	/*
		the string variable that we declare above has an underlying array of bytes which
		represents the utf-8 encoding of the full string which looks like following:

			r		  é				      s			u		  m		    é
		[01110010, 11000011, 10101001, 01110011, 01110101, 01101101, 11000011, 10101001]

		So, when we were indexing our string here what happened is that we're actually
		indexing the underlying byte array this is why we got 114 which is the value of resume[0]
		fmt.Println(resume[1])

		Note: if we index a string at the index of "é" we get 195 which is the first of the utf-8 encoding
		of this character, so we wouldnt get back the proper 233 we would expect for this character. But when
		we iterate over the string using the range keyword, we do get the 233 back, so the range keyword is doing
		some extra work for us here. It knows that our second character is 2 byte character and decodes it correctly
		to 233

		Take Away: When you're dealing with strings in Go, you are dealing with a value whose underlying representation is
		an array of bits. This is why taking the length of a string is length of bytes and not the number of characters.

		How to easily deal with iterating and indexing strings:
		Casting the string to an array of runes rather than dealing with underlying byte array of a string.firstChar

		var myString = []rune("Résumé")
		Example:
			Here what we get with our string when casting it to array of runes
			  R   é   s   u    m   é
			[114,233,155,117,109, 233]

		Runes:
			They are just unicode point numbers which represent the character.
			Runes are just an alias for int32 an dnow when we iterate we get the continuous index.\

		Note:
			we can declare a rune type using a single quote ' ' like so:
				var myRune = 'a'
	*/

	var resume1 = []rune("résumé")
	fmt.Println(resume1)
	var char = resume1[1]
	fmt.Printf("%v, %T\n", char, char) // 233, int32
	for i, v := range resume1 {
		fmt.Println(i, v)
		/*
			Output of the loop:
				index 	value
				  0		114
				  1		233
				  2		115
				  3		117
				  4		109
				  5		109
		*/
	}

	//declaring a rune type
	var myRune = 'a'
	fmt.Printf("\nMy Rune = %v of type %T", myRune, myRune)

	/*
		String building:
			we can concatenate strings using the plus symbols like so
			var string string = "a" + "b" or see code below

	*/

	var strSlice = []string{"t", "a", "s", "h", "i"}
	var name = ""
	for i := range strSlice {
		name += strSlice[i]
	}

	fmt.Printf("\n%v\n", name)

	//Note: Strings are immutable in GO meaning we cannot modify it once
	//created. Comment code below to see the error
	// name[0] = "T"

	/*
		so it follows that when we're concatenating a string and assigning it to a variable like this:
			name += strSlice[i]
		we're actually creating a completely new  string everytime which is pretty inefficient instead we can 
		import Go's built in package called strings and create a string builder, 
		instead of using a + operator we can use the WriteString() method and pass in the character we want to concatenate
		and at the end we call the String() method.
		
		What's Happening here:
			an array is allocated internally and values are appended when calling the writeString() method
			and only at the end a new string is created from the appended value when we call the String() method much faster
	
	*/

	var nameChars = []string{"T", "a", "s", "h", "i"}
	var strBuilder strings.Builder
	for i := range nameChars {
		strBuilder.WriteString(nameChars[i])
	}

	var myName = strBuilder.String()
	fmt.Println(myName)
}
