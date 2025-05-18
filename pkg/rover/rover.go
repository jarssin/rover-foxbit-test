package rover

import (
	"fmt"

	"github.com/jarssin/rover-foxbit-test/pkg/plateau"
)

type Direction string

type Rover struct {
	X, Y    int
	D       Direction
	plateau *plateau.Plateau
}

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

var leftTurn = map[Direction]Direction{
	North: West,
	West:  South,
	South: East,
	East:  North,
}

var rightTurn = map[Direction]Direction{
	North: East,
	East:  South,
	South: West,
	West:  North,
}

var moveDelta = map[Direction]struct{ dx, dy int }{
	North: {0, 1},
	South: {0, -1},
	East:  {1, 0},
	West:  {-1, 0},
}

func NewRover(x, y int, d Direction, p *plateau.Plateau) *Rover {
	return &Rover{
		X:       x,
		Y:       y,
		D:       d,
		plateau: p,
	}
}

func (r *Rover) Execute(commands string) error {
	for _, command := range commands {
		switch command {
		case 'L':
			r.D = leftTurn[r.D]
		case 'R':
			r.D = rightTurn[r.D]
		case 'M':
			if err := r.moveForward(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("invalid command: %c", command)
		}
	}
	return nil
}

func (r *Rover) moveForward() error {
	delta, ok := moveDelta[r.D]

	if !ok {
		return fmt.Errorf("invalid direction: %s", r.D)
	}

	newX := r.X + delta.dx
	newY := r.Y + delta.dy

	isOutOfBounds := newX < 0 || newY < 0 || newX > r.plateau.Width || newY > r.plateau.Height
	if isOutOfBounds {
		return fmt.Errorf("rover out of bounds")
	}

	r.X = newX
	r.Y = newY

	return nil
}
