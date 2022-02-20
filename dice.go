package main

import (
	"math/rand"
	"time"

	"github.com/ghfarrell/go-battleship/board"
	"github.com/ghfarrell/go-battleship/cprint"
)

/*

   ____________
  |            |
  |            |
  |            |
  |            |
  |____________|

*/
func RollTheBones() int {
	for i := 0; i < 2; i++ {
		board.ClearScreen()
		cprint.Printf["White"]("    _______\n")
		cprint.Printf["White"]("  /\\       \\\n")
		cprint.Printf["White"](" /()\\   ()  \\\n")
		cprint.Printf["White"]("/    \\_______\\\n")
		cprint.Printf["White"]("\\    /()     /\n")
		cprint.Printf["White"](" \\()/   ()  /\n")
		cprint.Printf["White"]("  \\/_____()/\n")
		time.Sleep(250 * time.Millisecond)
		board.ClearScreen()
		cprint.Printf["White"]("    _______\n")
		cprint.Printf["White"]("  /\\ ()  ()\\\n")
		cprint.Printf["White"](" /()\\   ()   \\\n")
		cprint.Printf["White"]("()   \\()___()\\\n")
		cprint.Printf["White"]("\\  ()/       /\n")
		cprint.Printf["White"](" \\()/   ()  /\n")
		cprint.Printf["White"]("  \\/_______/\n")
		time.Sleep(250 * time.Millisecond)
		board.ClearScreen()
		cprint.Printf["White"]("    _______\n")
		cprint.Printf["White"]("  /\\()   ()\\\n")
		cprint.Printf["White"](" /()\\()   ()\\\n")
		cprint.Printf["White"]("/()  \\()___()\\\n")
		cprint.Printf["White"]("\\  ()/()     /\n")
		cprint.Printf["White"](" \\()/       /\n")
		cprint.Printf["White"]("  \\/_____()/\n")
		time.Sleep(250 * time.Millisecond)
		board.ClearScreen()
		cprint.Printf["White"]("    _______\n")
		cprint.Printf["White"]("  /\\     ()\\\n")
		cprint.Printf["White"](" /()\\   ()  \\\n")
		cprint.Printf["White"]("/    \\()_____\\\n")
		cprint.Printf["White"]("\\    /()   ()/\n")
		cprint.Printf["White"](" \\()/   ()  /\n")
		cprint.Printf["White"]("  \\/()___()/\n")
		time.Sleep(250 * time.Millisecond)
	}
	r := rand.Intn(6) + 1
	resultDice(r)
	return r
}

func resultDice(r int) {
	switch r {
	case 1:
		board.ClearScreen()
		cprint.Printf["White"]("   ____________\n")
		cprint.Printf["White"]("  |            |\n")
		cprint.Printf["White"]("  |            |\n")
		cprint.Printf["White"]("  |     ()     |\n")
		cprint.Printf["White"]("  |            |\n")
		cprint.Printf["White"]("  |____________|\n")
		time.Sleep(1 * time.Second)
	case 2:
		board.ClearScreen()
		cprint.Printf["White"]("   ____________\n")
		cprint.Printf["White"]("  |            |\n")
		cprint.Printf["White"]("  |   ()       |\n")
		cprint.Printf["White"]("  |            |\n")
		cprint.Printf["White"]("  |        ()  |\n")
		cprint.Printf["White"]("  |____________|\n")
		time.Sleep(1 * time.Second)
	case 3:
		board.ClearScreen()
		cprint.Printf["White"]("   ____________\n")
		cprint.Printf["White"]("  |            |\n")
		cprint.Printf["White"]("  |  ()        |\n")
		cprint.Printf["White"]("  |     ()     |\n")
		cprint.Printf["White"]("  |        ()  |\n")
		cprint.Printf["White"]("  |____________|\n")
		time.Sleep(1 * time.Second)
	case 4:
		board.ClearScreen()
		cprint.Printf["White"]("   ____________\n")
		cprint.Printf["White"]("  |            |\n")
		cprint.Printf["White"]("  |  ()    ()  |\n")
		cprint.Printf["White"]("  |            |\n")
		cprint.Printf["White"]("  |  ()    ()  |\n")
		cprint.Printf["White"]("  |____________|\n")
		time.Sleep(1 * time.Second)
	case 5:
		board.ClearScreen()
		cprint.Printf["White"]("   ____________\n")
		cprint.Printf["White"]("  |            |\n")
		cprint.Printf["White"]("  |  ()    ()  |\n")
		cprint.Printf["White"]("  |     ()     |\n")
		cprint.Printf["White"]("  |  ()    ()  |\n")
		cprint.Printf["White"]("  |____________|\n")
		time.Sleep(1 * time.Second)
	case 6:
		board.ClearScreen()
		cprint.Printf["White"]("   ____________\n")
		cprint.Printf["White"]("  |            |\n")
		cprint.Printf["White"]("  |  ()    ()  |\n")
		cprint.Printf["White"]("  |  ()    ()  |\n")
		cprint.Printf["White"]("  |  ()    ()  |\n")
		cprint.Printf["White"]("  |____________|\n")
		time.Sleep(1 * time.Second)
	}
}
