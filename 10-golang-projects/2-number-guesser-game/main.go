package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"math/rand"
	"strconv"
)

func outputWelcomeMessage() {
	fmt.Println(strings.Repeat("-", 26))
	fmt.Println("Welcome to Number Guesser!")	
	fmt.Println(strings.Repeat("-", 26))
	fmt.Println()
}

func main () {
	reader := bufio.NewReader(os.Stdin)
	outputWelcomeMessage()

	fmt.Print("Enter max range for number to guess: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	handledInput, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		log.Fatal(err)
	}
	randomNumber := rand.Int() % handledInput + 1

	fmt.Println("\nNumber has been chosen. Guess away!\n")

	// main loop
	correctlyGuessed := false
	for correctlyGuessed != true {
		fmt.Print("Guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		handledInput, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatal(err)
		}

		if handledInput == randomNumber {
			correctlyGuessed = true
			fmt.Printf("That's it! %d was the number.\n", randomNumber)
		}
	}
}