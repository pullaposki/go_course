package main

import "fmt"

func main () {

	// *** Arrays ***

	// In go, arrays are fixed in size.
	// They are initiazlized to default values of their type.

	var myArray [6]int

	fmt.Println("My array ", myArray)
	fmt.Println("My array's length ", len(myArray))

	myArray[3] = 10
	fmt.Println("My array again",myArray)

	myInitializedArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Initialized array ", myInitializedArray)



	// *** Slices ***

	// This is a slice. Slices are dynamically created and can change in size.
	// They have more functionality than arrays.

	var mySlice []int // Does not allocate memory

	myAllocatedSlice := make([]int, 5) // Allocates memory

	fmt.Println("My slice ", mySlice)
	fmt.Println("My slice's length ", len(mySlice))

	fmt.Println("My allocated slice ", myAllocatedSlice)
	fmt.Println("My allocated slice's length ", len(myAllocatedSlice))

	// The append function is used to add elements to the end of the slice. The first argument to append is the slice you want to add elements to, which in this case is mySlice.

	mySlice = append(mySlice, 10)

	// Appending multiple values
	// []int {10,100}.... This is a slice literal (an array that we're defining right in the code) containing two integers, 10 and 100. 
	// The ... after the slice literal is a syntax feature of Go called "variadic syntax". It means "take this slice and unpack it into individual arguments". 	
	// So, instead of appending a single slice (which would nest one slice inside another), it appends each element of the slice individually.
	mySlice = append(mySlice, []int {10,100}...)

	fmt.Println("My slice after appending ", mySlice)
	fmt.Println("My slice's length after appending ", len(mySlice))

	copiedSlice := make([]int, len(mySlice))

	copy(copiedSlice, mySlice)
	fmt.Println("Copied slice ", copiedSlice)

	partialSlice := mySlice[1:3]
	fmt.Println("Partial slice ", partialSlice)


	

	// *** Maps ***
	// Maps can use any key value pair. They are like dictionaries.

	intStrMap := make(map[int]string)
	strIntMap := map[string]int{"One": 1, "Two": 2, "Three": 3}

	intStrMap[1] = "One"
	intStrMap[2] = "Two"
	intStrMap[3] = "Three"

	fmt.Println("IntStrMap ", intStrMap)
	fmt.Println("StrIntMap ", strIntMap)

	// Deleting from a map
	delete(intStrMap, 1)
	fmt.Println("IntStrMap after deleting ", intStrMap)
}