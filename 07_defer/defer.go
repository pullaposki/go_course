package main

import "fmt"

func willExecuteLast (greet string) {
	fmt.Printf("Goodbye %s, I was deferred to be last  in my calling function \n", greet)
}


func callsAdditionalDefer (greet string) {
	defer willExecuteLast(greet)
	fmt.Println("I will be before the first goodbye")
}

func helloGreeting( greet string) {
	fmt.Printf("Hello %s, I will execure first\n", greet)
}

// Using defer and panic to handle runtime errors
func panics() {
	panic("calamity ensues")
}

func main() {
	
	// if the function panics, the recover function will catch the panic, stop it from terminating the program, and return the value that was passed to panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("It panicked but we recovered. Error: %s\n", r)
		}
	}()

	defer panics()

	defer willExecuteLast("John")
	defer callsAdditionalDefer("Johhny")

	fmt.Println(("First we test defer"))
	helloGreeting("John")
}
