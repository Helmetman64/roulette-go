package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// array of colours
	var c [3]string
	c[0] = "Red"
	c[1] = "Black"
	c[2] = "Green"

	// counter to number the colours
	counter := 1

	// generates a random number from 0 to 100
	randomNumber := rand.Intn(100 - 0)
	// randomNumber := 100

	// empty variable of the chosen number
	cNumber := 0

	// empty variable of the chosen colour
	cColour := ""
	// empty variable of the winning colour
	wColour := ""

	fmt.Println("Welcome to Roulette! Choose either Black, Red or Green! If the colour you chose is the one displayed, you win!")

	fmt.Println("Choose a colour: ")
	fmt.Println("--------------------------------------------")

	// a for loop to iterate through the colour array
	for i := 0; i < len(c); i++ {

		// Displays the number to help order
		fmt.Printf("%v: ", counter)
		// Displays the colours from the colour array
		fmt.Println(c[i])
		counter++
	}
	fmt.Println("--------------------------------------------")

	fmt.Scan(&cNumber)

	/* if the chosen number is between 1 and 3, minus one to match the index of the colour array,
	then display the colour chosen and set the chosen colour to the value of the index in the colour array */
	if cNumber >= 1 && cNumber <= 3 {
		cNumber--
		fmt.Printf("You chose %v \n", c[cNumber])
		cColour = c[cNumber]
	} else {
		fmt.Println("Please choose a colour stated.")
		return
	}

	fmt.Println("--------------------------------------------")

	// if the random number is an even number and is not 0 or 100, set the winning colour to red
	if randomNumber%2 == 0 && randomNumber != 0 && randomNumber != 100 {
		fmt.Println("It's Red!")
		wColour = "Red"

		// if the random number is an odd number and is not 0 or 100, set the winning colour to black
	} else if randomNumber%2 != 0 && randomNumber != 0 && randomNumber != 100 {
		fmt.Println("It's Black!")
		wColour = "Black"

		// if the random number is 0 or 100, set the winning colour to green
	} else {
		fmt.Println("It's Green!")
		wColour = "Green"
	}

	// if the chosen colour matches the winning colour, the player wins
	if cColour == wColour {
		fmt.Println("You Win!")
	} else {
		fmt.Println("You Lose!")
	}

}
