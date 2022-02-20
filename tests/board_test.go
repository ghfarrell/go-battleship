package board_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/ghfarrell/go-battleship/board"
	"github.com/ghfarrell/go-battleship/cprint"
)

var b board.Board

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
	b.Board = make(map[board.Point]byte, 100)
	for c := 'a'; c <= 'j'; c++ {
		for i := 1; i <= 10; i++ {
			b.Board[board.CoordToPoint(c, i)] = '-'
		}
	}
	b.Ships = make([]board.Ship, 5)
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
		s.Coords = make([]board.Point, 0)
	}
}
func TestPlaceShip(t *testing.T) {
	s := b.Ships[0]
	testPoint := board.Point{
		Row: 'a',
		Col: 1,
	}
	solPoints := make([]board.Point, 0)
	solPoints = append(solPoints, board.Point{Row: 'a', Col: 1})
	solPoints = append(solPoints, board.Point{Row: 'b', Col: 1})
	solPoints = append(solPoints, board.Point{Row: 'c', Col: 1})
	solPoints = append(solPoints, board.Point{Row: 'd', Col: 1})
	solPoints = append(solPoints, board.Point{Row: 'e', Col: 1})
	s, _ = b.PlaceShip(s, testPoint, true)
	cprint.Printf["Red"]("%s: %v\n", s.Name, s.Coords)
	for i := 0; i < 5; i++ {
		cprint.Printf["Red"]("%v\n", s.Coords[i])
		if s.Coords[i] != solPoints[i] {
			t.Errorf("Wrong points; Expected %v got %v", solPoints, s.Coords)
		}
	}
}

func TestAutoPlaceShips(t *testing.T) {
	b.AutoPlaceShips()
	for _, s := range b.Ships {
		if len(s.Coords) < 1 {
			t.Errorf("Expect coordinate array, got %v", s.Coords)
		}
	}
}
