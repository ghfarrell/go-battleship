/*
	TODO:
	- give AI a hard mode
	- make AI communicate with player (sassy)
*/

/*
	COLORS:
	- Cyan: Game talking to user
	- Red: Errors or hit
	- White/Yellow: Expecting input
	- Green: Response to input
	- Magenta: Static game text and game over text
	- Blue: Debug and AI talking
*/

package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/ghfarrell/go-battleship/board"
	"github.com/ghfarrell/go-battleship/cprint"
	"github.com/ghfarrell/go-battleship/player"
)

func PlayerGoesFirst() bool {
	var playerGuess int
	for {
		cprint.Printf["Cyan"]("Time to determine who goes first! Guess a number 1-6: ")
		_, err := fmt.Scanf("%d\n", &playerGuess)
		if err != nil {
			cprint.Printf["Red"]("Bad input! Try again\n")
		} else {
			break
		}
	}
	aiGuess := rand.Intn(6) + 1
	cprint.Printf["Cyan"]("The AI made its guess! Time to see who is closest...\n")
	time.Sleep(1 * time.Second)
	r := RollTheBones()
	fmt.Println()
	cprint.Printf["Cyan"]("The number is %d!\n", r)
	cprint.Printf["Cyan"]("The AI guessed %d, so the winner is...\n", aiGuess)
	if math.Abs(float64(r-playerGuess)) < math.Abs(float64(r-aiGuess)) {
		cprint.Printf["Cyan"]("You! You will make your strike first.\n")
	} else if math.Abs(float64(r-playerGuess)) == math.Abs(float64(r-aiGuess)) {
		cprint.Printf["Cyan"]("You... via tiebreaker!\nYou both guessed %d, but I like you more than the AI.\n", playerGuess)
	} else {
		cprint.Printf["Cyan"]("The AI! Too bad! They will strike first.\n")
	}
	time.Sleep(3 * time.Second)
	return true
}

func playerWin() {

}

func aiWin() {

}

func main() {
	var (
		User      player.Player
		ai        player.AI
		userFirst bool
	)
	debug := flag.Bool("debug", false, "Turns debug mode on")
	flag.Parse()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cprint.Printf["Red"]("\nExiting Battleship... ")
		os.Exit(1)
	}()
	User.Initialize()
	ai.Initialize()
	if *debug {
		User.DebugPlaceShips()
		cprint.Printf["Magenta"]("User ships placed...\n")
		ai.PlaceShips()
		cprint.Printf["Magenta"]("AI ships placed...\n")
		userFirst = true
	} else {
		board.ClearScreen()
		cprint.Printf["Cyan"]("Time to place your ships!\n\n")
		User.PlaceShips()
		ai.PlaceShips()
		userFirst = PlayerGoesFirst()
	}
	//main gameplay loop
	if userFirst {
		for {
			if User.Guess(&ai, *debug) || ai.Guess(&User, *debug) {
				break
			}

		}
	} else {
		for {
			if ai.Guess(&User, *debug) || User.Guess(&ai, *debug) {
				break
			}
		}
	}
}
