package main

import (
	"fmt"
	"math/rand"
	"time"
)

func opponent_worker(messages chan int) {
	// Randomize the choice between rock, paper, and scissors
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	choice := random.Intn(3)

	// Send the choice to the main thread
	messages <- choice
}

func main() {
	// create the channels
	messages := make(chan int)
	done := make(chan bool)

	winPoints := 3

	playerPoints := 0
	opponentPoints := 0

	fmt.Println("Welcome to Rock, Paper, Scissors!")

	// loop until either the player or the opponent reaches the winPoints
	for playerPoints < winPoints && opponentPoints < winPoints {
		go func() {
			var _ chan bool = done
			opponent_worker(messages)
		}()

		// get the opponent's choice from the opponent worker
		opponentChoice := <-messages

		// get the user's choice
		var userChoice int
		fmt.Println("Enter your choice (0 for rock, 1 for paper, 2 for scissors): ")

		for {
			
			// use blank identifier to get the error value
			_, err := fmt.Scanln(&userChoice)

			// check if the input is valid
			if err == nil && (userChoice == 0 || userChoice == 1 || userChoice == 2) {
				break
			}
			fmt.Println("Invalid input. Please enter 0 for rock, 1 for paper, or 2 for scissors: ")
		}

		printUserChoice(userChoice)
		printOpponentChoice(opponentChoice)

		switch {
		case userChoice == opponentChoice:
			fmt.Println("It's a tie!")
		case userChoice == 0 && opponentChoice == 1:
			fmt.Println("Opponent wins!")
			opponentPoints++
		case userChoice == 0 && opponentChoice == 2:
			fmt.Println("You win!")
			playerPoints++
		case userChoice == 1 && opponentChoice == 0:
			fmt.Println("You win!")
			playerPoints++
		case userChoice == 1 && opponentChoice == 2:
			fmt.Println("Opponent wins!")
			opponentPoints++
		case userChoice == 2 && opponentChoice == 0:
			fmt.Println("Opponent wins!")
			opponentPoints++
		}

		fmt.Println("Player Points: ", playerPoints)
		fmt.Println("Opponent Points: ", opponentPoints)

		if playerPoints == winPoints {
			fmt.Println("You win the game!")
		} else if opponentPoints == winPoints {
			fmt.Println("Opponent wins the game!")
		}

		// wait before starting the next round
		time.Sleep(1 * time.Second)
	}
}

func printUserChoice(userChoice int) {
	switch userChoice {
	case 0:
		fmt.Println("You chose rock")
	case 1:
		fmt.Println("You chose paper")
	case 2:
		fmt.Println("You chose scissors")
	}
}

func printOpponentChoice(opponentChoice int) {
	switch opponentChoice {
	case 0:
		fmt.Println("Opponent chose rock")
	case 1:
		fmt.Println("Opponent chose paper")
	case 2:
		fmt.Println("Opponent chose scissors")
	}
}
