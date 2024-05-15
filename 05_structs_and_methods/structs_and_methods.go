package main

import "fmt"

type dog struct {
	name, breed string
}

type cat struct {
	name, breed string
}

// Interfaces are implemented by structs by implementing each function separately.

type animal interface {
	move()
	speak()
}

// The (d dog) and (c cat) are called receivers.
// The receiver is a parameter of a method that is called on a type.
// The receiver is used to access the properties of the type.
// The receiver is used to call the methods of the type.

func (d dog) move() {
	fmt.Println(d.name, "is running")
}

func (d dog) speak() {
	fmt.Println(d.name, "is barking")
}

func (c cat) move() {
	fmt.Println(c.name, "is running")
}

func (c cat) speak() {
	fmt.Println(c.name, "is meowing")
}

func act (a animal) {
	a.move()
	a.speak()
}

func main() {
	duke := dog{name: "Duke", breed: "German Shepherd"}
	whiskers := cat{name: "Whiskers", breed: "Siamese"}

	fmt.Println("Accessing directly through methods")

	duke.move()
	duke.speak()
	whiskers.move()
	whiskers.speak()

	fmt.Println("\nAccessing through interface")

	act(duke)
	act(whiskers)
}