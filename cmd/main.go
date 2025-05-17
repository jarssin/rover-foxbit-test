package main

import (
	"fmt"
	"rover/internal/utils"
	"rover/pkg/plateau"
	"rover/pkg/rover"
)

func main() {
	reader := utils.NewInputReader()
	plateauWidth, plateauHeight, err := reader.PlateauSize()

	if err != nil {
		return
	}

	p := &plateau.Plateau{
		Width:  plateauWidth,
		Height: plateauHeight,
	}

	for {
		x, y, d, err := reader.RoverPosition()

		if err != nil {
			return
		}

		rover := rover.NewRover(x, y, d, p)
		fmt.Printf("Rover: %+v\n", rover)

		// commands, err := reader.RoverCommands()
		// if err != nil {
		// 	return
		// }

	}
}
