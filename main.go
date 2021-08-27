package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func generateRandomInt() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(100)
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

func main() {
}
