package board

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode"

	"github.com/ghfarrell/go-battleship/cprint"
)

type Point struct {
	Row rune
	Col int
}

type Ship struct {
	Name   string
	Length int
	Coords []Point
}

/*
	Enemy board's state is - for an empty space, O for a missed
	shot, and X for a hit
	Friendly board's state is - for an empty space, O for a boat
	and X for a boat space that has been hit
*/
type Board struct {
	Board map[Point]byte
	Ships []Ship
}

func (b *Board) Initialize() {
	rand.Seed(int64(time.Now().Nanosecond()))
	b.Board = make(map[Point]byte, 100)
	for c := 'a'; c <= 'j'; c++ {
		for i := 1; i <= 10; i++ {
			b.Board[CoordToPoint(c, i)] = '-'
		}
	}
	b.Ships = make([]Ship, 5)
	b.Ships[0].Name = "Carrier"
	b.Ships[1].Name = "Battleship"
	b.Ships[2].Name = "Cruiser"
	b.Ships[3].Name = "Submarine"
	b.Ships[4].Name = "Destroyer"
	b.Ships[0].Length = 5
	b.Ships[1].Length = 4
	b.Ships[2].Length = 3
	b.Ships[3].Length = 3
	b.Ships[4].Length = 2
	for _, s := range b.Ships {
		s.Coords = make([]Point, 0)
	}
}

/*
	helper func turns 0 to false and anything else to true
	for ai placing ship logic
*/
func itob(i int) bool {
	if i == 0 {
		return false
	} else {
		return true
	}
}

/*
	helper func turns "v" or "h" to true and false respectively
	for the ship placement logic
*/
func HVToBool(c rune) bool {
	if c == 'h' || c == 'H' {
		return false
	} else if c == 'v' || c == 'V' {
		return true
	} else {
		cprint.Printf["Red"]("Invalid orientation! Defaulting to vertical\n")
		return true
	}
}

func RandPoint() Point {
	rows := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	r := rows[rand.Intn(10)]
	c := rand.Intn(10) + 1
	return Point{r, c}
}

func InputPoint() (Point, error) {
	var row rune
	var col int
	_, err := fmt.Scanf("%c%d\n", &row, &col)
	if err != nil {
		cprint.Printf["Red"]("Bad input. Enter coordinate as [row][col] i.e. b5\n")
	}
	if unicode.IsUpper(row) {
		row = unicode.ToLower(row)
	}
	if col < 1 || col > 10 || row < 'a' || row > 'j' {
		return Point{}, fmt.Errorf("Invalid coordinate")
	} else {
		return Point{row, col}, nil
	}
}

func (p Point) IsValid() bool {
	if p.Row > 'j' || p.Row < 'a' || p.Col > 10 || p.Col < 1 {
		return false
	}
	return true
}

func (p Point) NudgeRight(i int) Point {
	return Point{p.Row, p.Col + i}
}
func (p Point) NudgeUp(i int) Point {
	return Point{p.Row - rune(i), p.Col}
}
func (p Point) NudgeDown(i int) Point {
	return Point{p.Row + rune(i), p.Col}
}
func (p Point) NudgeLeft(i int) Point {
	return Point{p.Row, p.Col - i}
}

func CoordToPoint(b rune, i int) Point {
	if unicode.IsUpper(b) {
		b = unicode.ToLower(b)
	}
	return Point{b, i}
}

func (b Board) print() {
	fmt.Println("  1 2 3 4 5 6 7 8 9 10")
	for c := 'a'; c <= 'j'; c++ {
		fmt.Printf("%c", c)
		for i := 1; i <= 10; i++ {
			cprint.Printf["White"](" " + string(b.Board[CoordToPoint(c, i)]))
		}
		fmt.Println()
	}
}

func (b Board) GameOver() bool {
	for c := 'a'; c <= 'j'; c++ {
		for i := 1; i <= 10; i++ {
			if b.Board[CoordToPoint(c, i)] == 'O' {
				return false
			}
		}
	}
	return true
}

func (b Board) shipExistsAt(coords []Point) bool {
	for _, p := range coords {
		if b.Board[p] != '-' {
			return true
		}
	}
	return false
}

// ok so for some stupid reason that i cant figure out, even if i pass the ship
// parameter here by reference the new coordinate array that the ship contains
// WILL NOT be updated when the function returns, and i literally have no idea why
// so thats why it returns the Ship with the new array as well as the error
func (b *Board) PlaceShip(s Ship, p Point, vertical bool) (Ship, error) {
	switch vertical {
	case true:
		if p.Row+rune(s.Length-1) > 'j' {
			break
		}
		for i := 0; i < s.Length; i++ {
			s.Coords = append(s.Coords, Point{p.Row + rune(i), p.Col})
		}
		if !b.shipExistsAt(s.Coords) {
			for _, c := range s.Coords {
				b.Board[c] = 'O'
			}
		} else {
			s.Coords = make([]Point, 0)
			return s, fmt.Errorf("Ship already exists at that point")
		}
		return s, nil
	case false:
		if p.Col+s.Length-1 > 10 {
			break
		}
		for i := 0; i < s.Length; i++ {
			s.Coords = append(s.Coords, Point{p.Row, p.Col + i})
		}
		if !b.shipExistsAt(s.Coords) {
			for _, c := range s.Coords {
				b.Board[c] = 'O'
			}
		} else {
			s.Coords = make([]Point, 0)
			return s, fmt.Errorf("Ship already exists at that point")
		}
		return s, nil
	default:

		s.Coords = make([]Point, 0)
		return s, fmt.Errorf("Coordinate out of bounds")
	}

	s.Coords = make([]Point, 0)
	return s, fmt.Errorf("Coordinate out of bounds.")
}

func (b *Board) PlaceShips() {
	stdin := bufio.NewReader(os.Stdin)
	var vert rune
	var vertBool bool
	for i := range b.Ships {
		s := b.Ships[i]
		for {
			b.print()
			cprint.Printf["Yellow"]("Where do you want your %s? (Length %d): ", s.Name, s.Length)
			p, err := InputPoint()
			if err != nil {
				cprint.Printf["Red"]("Error: " + err.Error() + "\n")
			}
			cprint.Printf["White"]("Placed Horizontally or Vertically (h or v): ")
			_, err = fmt.Scanf("%c", &vert)
			stdin.ReadString('\n')
			if err != nil {
				stdin.ReadString('\n')
				cprint.Printf["Red"]("Bad input! Enter 'v', 'h', 'vertical', or 'horizontal'\n")
			}
			vertBool = HVToBool(vert)
			if b.Ships[i], err = b.PlaceShip(s, p, vertBool); err == nil {
				s := b.Ships[i]
				ClearScreen()
				if vertBool {
					cprint.Printf["Cyan"]("%s placed vertically at coordinate [%c%d]!\n", s.Name, p.Row, p.Col)
				} else {
					cprint.Printf["Cyan"]("%s placed horizontally at coordinate [%c%d]!\n", s.Name, p.Row, p.Col)
				}

				break
			} else {
				cprint.Printf["Red"]("Error: " + err.Error() + "\n")
			}
		}
	}
}

func (b *Board) AutoPlaceShips() {
	var err error
	for i := range b.Ships {
		for {
			s := b.Ships[i]
			p := RandPoint()
			v := rand.Intn(2)
			b.Ships[i], err = b.PlaceShip(s, p, itob(v))
			s = b.Ships[i]
			if err == nil {
				break
			}
		}
	}
}

func (b Board) Sunk(s Ship) bool {
	for _, p := range s.Coords {
		if b.Board[p] == 'O' { //there are still spots yet to be hit on the ship
			time.Sleep(5)
			return false
		}
	}
	return true
}

/*
	returns whether or not the shot hit, and if the ship has been sunk,
	it returns the sunken ship. otherwise, it returns an empty Ship
*/
func (b Board) CheckForHit(c Point) (hit bool, ship Ship) {
	if b.Board[c] == 'O' {
		hit = true
		for _, s := range b.Ships {
			for _, p := range s.Coords {
				if p == c {
					ship = s
					break
				}
			}
			if ship.Name == s.Name {
				break
			}
		}
	} else {
		hit = false
		ship = Ship{}
	}
	return hit, ship
}
