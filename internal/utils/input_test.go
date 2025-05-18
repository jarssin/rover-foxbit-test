package utils

import (
	"strings"
	"testing"

	"github.com/jarssin/rover-foxbit-test/pkg/rover"
)

func TestPlateauSize(t *testing.T) {
	input := "5 5\n"
	r := NewInputReaderFrom(strings.NewReader(input))
	w, h, err := r.PlateauSize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w != 5 || h != 5 {
		t.Errorf("expected 5 5, got %d %d", w, h)
	}
}

func TestRoverPosition(t *testing.T) {
	input := "1 2 N\n"
	r := NewInputReaderFrom(strings.NewReader(input))
	x, y, d, err := r.RoverPosition()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if x != 1 || y != 2 || d != rover.North {
		t.Errorf("expected 1 2 N, got %d %d %s", x, y, d)
	}
}

func TestRoverCommands(t *testing.T) {
	input := "LMLMLMLMM\n"
	r := NewInputReaderFrom(strings.NewReader(input))
	cmds, err := r.RoverCommands()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cmds != "LMLMLMLMM" {
		t.Errorf("expected LMLMLMLMM, got %s", cmds)
	}
}
