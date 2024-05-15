package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func math_action(a int, b int, action string) (int, string) {
	if action == "add" {
		return add(a, b), "add"
	} else if action == "sub" {
		return sub(a, b), "sub"
	} else {
		return 0, "error"
	}
}

// Variadic functions
func sum(vals ...int) int {
	total := 0
	for _, num := range vals {
		total += num
	}
	return total
}

// function return values
func sequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	c := add(20, 10)
	d := sub(20, 10)
	fmt.Println("20 + 10 = ", c)
	fmt.Println("20 - 10 = ", d)

	e, act := math_action(20, 10, "add")
	f, act2 := math_action(20, 10, "sub")
	fmt.Println("20", act, "10 = ", e)
	fmt.Println("20 - 10 = ", f, act2)

	t := []	int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Adding together these numbers: ", t)
	fmt.Printf("leads to %d\n", sum(t...))

	nextInt := sequence()
	fmt.Println("ints in sequence: ")
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	moreInts := sequence()
	fmt.Println("Begin again")
	fmt.Println(moreInts())
	fmt.Println(moreInts())
	fmt.Println(moreInts())

	println(math_action(5, 3, "add"))
	println(math_action(5, 3, "sub"))
	println(math_action(5, 3, "error"))
}