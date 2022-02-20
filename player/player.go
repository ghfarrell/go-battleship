package player

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ghfarrell/go-battleship/board"
	"github.com/ghfarrell/go-battleship/cprint"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) opposite() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	default:
		//impossible but compiler yells at me if i dont
		return Up
	}
}

/*
	In ship slice, ships are ordered by length, longest to smallest, and alphabetically
*/
type Player struct {
	Guesses  []board.Point
	Friendly board.Board
	Enemy    board.Board
}

type AI struct {
	// a Column value of -1 means the pivot is not in use
	Pivot       board.Point
	PivotOffset int
	Dir         Direction
	Player
}

func (p *Player) Initialize() {
	p.Friendly.Initialize()
	p.Enemy.Initialize()
	p.Guesses = make([]board.Point, 0)
}
func (p *AI) Initialize() {
	p.Friendly.Initialize()
	p.Enemy.Initialize()
	p.Guesses = make([]board.Point, 0)
	p.PivotOffset = 1
	p.Pivot = board.Point{
		Row: 'a',
		Col: -1,
	}
	p.Dir = Up
}

func (p Player) PrintFriendly() {
	fmt.Println("  1 2 3 4 5 6 7 8 9 10")
	for c := 'a'; c <= 'j'; c++ {
		fmt.Printf("%c", c)
		for i := 1; i <= 10; i++ {
			cprint.Printf["White"](" " + string(p.Friendly.Board[board.CoordToPoint(c, i)]))
		}
		fmt.Println()
	}
}
func (p Player) PrintEnemy() {
	fmt.Println("  1 2 3 4 5 6 7 8 9 10")
	for c := 'a'; c <= 'j'; c++ {
		fmt.Printf("%c", c)
		for i := 1; i <= 10; i++ {
			cprint.Printf["Red"](" " + string(p.Enemy.Board[board.CoordToPoint(c, i)]))
		}
		fmt.Println()
	}
}

func hitResponse() string {
	r := []string{
		"A perfect shot!",
		"Nice hit!",
		"Dead on!",
		"Incredible shot!",
		"Nice one!",
	}
	return r[rand.Intn(len(r))]
}
func missResponse() string {
	r := []string{
		"So close!",
		"Try hitting a boat next time...",
		"Almost!",
		"Too bad!",
		"No dice!",
	}
	return r[rand.Intn(len(r))]
}

func (p Player) hasGuessed(c board.Point) bool {
	for _, i := range p.Guesses {
		if i == c {
			return true
		}
	}
	return false
}

func (p Player) DebugPrintGuesses() {
	cprint.Printf["Blue"]("[")
	for _, c := range p.Guesses {
		cprint.Printf["Blue"](" %c%d ", c.Row, c.Col)
	}
	cprint.Printf["Blue"]("]")
}

func (p *Player) PlaceShips() {
	p.Friendly.PlaceShips()
}
func (p *Player) DebugPlaceShips() {
	p.Friendly.AutoPlaceShips()
}

func (a *AI) PlaceShips() {
	a.Friendly.AutoPlaceShips()
}

func (p *Player) Guess(a *AI, debug bool) bool {
	var (
		c   board.Point
		err error
	)
	board.ClearScreen()
	p.PrintEnemy()
	p.PrintFriendly()
	if debug {
		a.PrintFriendly()
	}
	if debug {
		cprint.Printf["Blue"]("DEBUG: USER GUESSES -> ")
		p.DebugPrintGuesses()
		cprint.Printf["Blue"]("\nDEBUG: AI GUESSES -> ")
		a.DebugPrintGuesses()
		fmt.Println()
	}
	for {
		cprint.Printf["Yellow"]("Pick a coordinate to attack: ")
		c, err = board.InputPoint()
		if err != nil {
			cprint.Printf["Red"]("Error: " + err.Error() + "\n")
		} else if p.hasGuessed(c) {
			cprint.Printf["Red"]("Woops! You already guessed that spot.\n")
		} else {
			p.Guesses = append(p.Guesses, c)
			break
		}
	}
	if hit, ship := a.Friendly.CheckForHit(c); hit {
		a.Friendly.Board[c] = 'X'
		p.Enemy.Board[c] = 'X'
		cprint.Printf["Green"](hitResponse() + "\n")
		cprint.Printf["Green"]("You hit the AI's %s!\n", ship.Name)
		if a.Friendly.Sunk(ship) {
			cprint.Printf["Green"]("The AI's %s sunk!\n", ship.Name)
		}
		if a.Friendly.GameOver() {
			cprint.Printf["Magenta"]("YOU WIN! THE EVIL AI IS DEFEATED!\n")
			return true
		}
	} else {
		p.Enemy.Board[c] = 'O'
		cprint.Printf["Red"](missResponse() + "\n")
	}
	time.Sleep(2 * time.Second)
	return false
}

func (a AI) randomGuess() board.Point {
	for {
		c := board.RandPoint()
		if !a.hasGuessed(c) {
			return c
		}
	}
}

/*
	guessing alg should work like:
	- randomly guess a spot until you hit, that spot is the pivot
	- go one direction (say, up) until either you miss or you sink
	- if you miss, go one spot below the pivot
	- if that misses, go left from the pivot
	- if that misses, go right from the pivot
	- since you go all directions, one must work and you sink a ship
	- when you sink, go back to random guesses and repeat
*/

/*
	TODO:
	- it guessed up and missed, so when it guessed right and hit, it guessed that
		exact same spot again which is impossible to guess the same spot twice...
*/
func (a *AI) getGuess() (c board.Point) {
	if a.Pivot.Col == -1 || len(a.Guesses) < 1 {
		c = a.randomGuess()
		return
	}
	switch a.Dir {
	case Up:
		c = a.Pivot.NudgeUp(a.PivotOffset)
	case Down:
		c = a.Pivot.NudgeDown(a.PivotOffset)
	case Left:
		c = a.Pivot.NudgeLeft(a.PivotOffset)
	case Right:
		c = a.Pivot.NudgeRight(a.PivotOffset)
	}
	if a.hasGuessed(c) { // if the bot somehow comes up with a point that has already been guessed,
		c = a.randomGuess() // ill deal with it later
	} else if c.IsValid() {
		a.PivotOffset++
		return
	} else { // if the nudged coordinate is out of bounds, go the opposite way
		a.PivotOffset = 1
		a.Dir = a.Dir.opposite()
		c = a.getGuess()
	}
	return
}

func (a *AI) Guess(p *Player, debug bool) bool {
	c := a.getGuess()
	a.Guesses = append(a.Guesses, c)
	cprint.Printf["White"]("Now the AI's time to strike...\n")
	if !debug {
		time.Sleep(1 * time.Second)
	}
	cprint.Printf["White"]("The AI has fired at coordinate %c%d!\n", c.Row, c.Col)
	if !debug {
		time.Sleep(1 * time.Second)
	}
	if hit, sunkShip := p.Friendly.CheckForHit(c); hit {
		cprint.Printf["Red"]("You've been hit!\n")
		p.Friendly.Board[c] = 'X'
		if a.Pivot.Col == -1 { // if AI wasn't guessing on a pivot and hit,
			a.Pivot = c // set the new pivot
			a.PivotOffset = 1
			a.Dir = Up
			if debug {
				cprint.Printf["Blue"]("DEBUG: New pivot: %c%d\n", a.Pivot.Row, a.Pivot.Col)
			}
		} // else, keep on guessing on that pivot
		cprint.Printf["Red"]("Your %s has taken a hit!\n", sunkShip.Name)
		if p.Friendly.Sunk(sunkShip) {
			// a little more complex case: if the AI was guessing on a line and sunk a ship that is
			// shorter than the line it was guessing on, then the line must be composed of multiple
			// ships, so after sinking one of the two ships, it will travel in the opposite direction
			// to sink the other ship that composed that line
			if a.PivotOffset > sunkShip.Length {
				a.PivotOffset = 1
				a.Dir = a.Dir.opposite()
			} else {
				// reset pivot
				a.Pivot.Col = -1
				cprint.Printf["Red"]("The AI has sunk your %s!\n", sunkShip.Name)
			}
		}
		if p.Friendly.GameOver() {
			cprint.Printf["Magenta"]("OH GLOB! YOU LOST! YOU STINK!\n")
			return true

		}
	} else {
		if a.Pivot.Col != -1 { // if AI was guessing on a pivot and missed...
			if a.PivotOffset > 2 { // if the AI was following a track of hits and missed
				a.Dir = a.Dir.opposite() // go the opposite direction
			} else { // cycle directions
				a.Dir++
			}
			a.PivotOffset = 1
		}
		cprint.Printf["Green"]("Phew! The AI doesn't have very good aim.")
	}
	time.Sleep(3 * time.Second)
	return false
}
