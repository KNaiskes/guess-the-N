package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func generateRandomInt(rangeInt int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(rangeInt)
}

func getUserInt() int {
	var userInput string
	var convertToInt int

	for {
		_, err := fmt.Scan(&userInput)
		convertToInt, err = strconv.Atoi(userInput)
		if err != nil {
			// fmt.Println(err)
			fmt.Println("Enter a valid integer: ")
		} else {
			break
		}
	}
	return convertToInt
}

func correctGuess(userGuess int, randomGuess int) bool {
	if userGuess == randomGuess {
		return true
	} else {
		return false
	}
}

func main() {
	const MAX = 10

	randomGuess := generateRandomInt(MAX)
	// fmt.Printf("Random number %d\n", randomGuess)

	fmt.Printf("Make a guess, the bigger number is guaranteed to be: %d\n", MAX)
	userGuess := getUserInt()

	for {
		if correctGuess(userGuess, randomGuess) {
			fmt.Println("WOW! You guessed correctly")
			break
		} else {
			fmt.Println("Your guess is wrong :(")
			fmt.Println("Try again")
			userGuess = getUserInt()
		}
	}
}
