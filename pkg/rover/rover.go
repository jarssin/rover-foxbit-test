package rover

import (
	"rover/pkg/plateau"
)

type RoverInterface interface{}

type Direction string

type Rover struct {
	x, y    int
	d       Direction
	plateau *plateau.Plateau
}

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

func NewRover(x, y int, d Direction, p *plateau.Plateau) RoverInterface {
	return &Rover{
		x:       x,
		y:       y,
		d:       d,
		plateau: p,
	}
}
