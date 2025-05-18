package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/jarssin/rover-foxbit-test/pkg/rover"
)

type InputReaderInterface interface {
	PlateauSize() (int, int, error)
	RoverPosition() (int, int, rover.Direction, error)
	RoverCommands() (string, error)
}

type InputReader struct {
	reader *bufio.Reader
}

func NewInputReader() InputReaderInterface {
	return &InputReader{reader: bufio.NewReader(os.Stdin)}
}

func NewInputReaderFrom(r io.Reader) InputReaderInterface {
	return &InputReader{reader: bufio.NewReader(r)}
}

func (ir *InputReader) PlateauSize() (int, int, error) {
	text, err := ir.reader.ReadString('\n')
	if err != nil {
		return 0, 0, err
	}
	parts := strings.Fields(text)
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid input format")
	}
	width, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	height, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}
	return height, width, nil
}

func (ir *InputReader) RoverPosition() (int, int, rover.Direction, error) {
	text, err := ir.reader.ReadString('\n')
	if err != nil {
		return 0, 0, "", err
	}

	parts := strings.Fields(text)
	if len(parts) != 3 {
		return 0, 0, "", fmt.Errorf("invalid input format")
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, "", err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, "", err
	}

	direction := rover.Direction(parts[2])
	validDirections := map[rover.Direction]bool{
		rover.North: true,
		rover.East:  true,
		rover.South: true,
		rover.West:  true,
	}
	if !validDirections[direction] {
		return 0, 0, "", fmt.Errorf("invalid direction")
	}

	return x, y, direction, nil
}

func (ir *InputReader) RoverCommands() (string, error) {
	text, err := ir.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	commands := strings.TrimSpace(text)
	if len(commands) == 0 {
		return "", fmt.Errorf("no commands provided")
	}

	return commands, nil
}
