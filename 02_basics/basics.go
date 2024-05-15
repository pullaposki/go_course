package main

import "fmt"

func main() {
	// main is the entry point of the Go program.
	// It initializes variables, demonstrates various control flow statements,
	// and calls the `whatAmI` function with different types of arguments.
	var b int = 15
	var a int
	var t int = 10

	// Initializing array to size  6. No need to give initial values to all indexes
	// Uninitialized indexes will be set to 0
	numbers := [6]int{1, 2, 3, 5}

	for a := 0; a < 10; a++ {
		fmt.Printf("Value of a: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("Value of a: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("Value of x = %d at %d\n", x, i)}

		if t > 5 {
			fmt.Printf("t is greater than 5\n")
		}else {
			fmt.Printf("t is less than 5\n")
		}

		j := 2
		fmt.Print("Write ", j, " as")
		switch j {
		case 1:
			fmt.Println(" one")
		case 2:
			fmt.Println(" two")
		case 3:
			fmt.Println(" three")
		}

		// whatAmI is a function that takes an interface{} as a parameter and determines its type.
		// It prints a message indicating the type of the parameter.
		whatAmI := func(i interface{}) {
			switch t := i.(type) {
			case bool:
				fmt.Println("I'm a bool")
			case int:
				fmt.Println("I'm an int")
			default:
				fmt.Printf("Don't know type %T\n", t)
			}
		}

		whatAmI(true)
		whatAmI(1)
		whatAmI("hey")

	}