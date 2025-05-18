package main

import (
	"fmt"

	"github.com/jarssin/rover-foxbit-test/internal/utils"
	"github.com/jarssin/rover-foxbit-test/pkg/plateau"
	"github.com/jarssin/rover-foxbit-test/pkg/rover"
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

		commands, err := reader.RoverCommands()
		if err != nil {
			return
		}

		if err := rover.Execute(commands); err != nil {
			fmt.Println("Error executing commands:", err)
			return
		}

		fmt.Printf("%d %d %s\n", rover.X, rover.Y, rover.D)
	}
}
