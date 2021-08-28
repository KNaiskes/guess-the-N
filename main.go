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
	const TRIES = 2

	fmt.Println("Random number")
	randomGuess := generateRandomInt(10)
	fmt.Println(randomGuess)

	fmt.Println("Make a guess")
	userGuess := getUserInt()

	// fmt.Println(userGuess)

	for i := 0; i < TRIES; i++ {
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
