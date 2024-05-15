package main

import (
	"errors"
	"fmt"
)

	func notDog(animal string) (string, error) {
		if animal != "dog" {
			return "", errors.New("Not a dog")
		}
		return "It's a dog", nil
	}


	func main() {
		animals := []string {"cat", "dog", "fish"}

		for _, i := range animals  {
			result, err := notDog(i)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println(result)
			}
		}
	}