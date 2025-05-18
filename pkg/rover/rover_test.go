package rover

import (
	"testing"

	"github.com/jarssin/rover-foxbit-test/pkg/plateau"
)

func TestRover_Execute_Success(t *testing.T) {
	p := &plateau.Plateau{Width: 5, Height: 5}
	r := NewRover(1, 2, North, p)
	err := r.Execute("LMLMLMLMM")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.X != 1 || r.Y != 3 || r.D != North {
		t.Errorf("expected 1 3 N, got %d %d %s", r.X, r.Y, r.D)
	}
}

func TestRover_Execute_InvalidCommand(t *testing.T) {
	p := &plateau.Plateau{Width: 5, Height: 5}
	r := NewRover(1, 2, North, p)
	err := r.Execute("LMLX")
	if err == nil {
		t.Error("expected error for invalid command, got nil")
	}
}

func TestRover_Move_OutOfBounds(t *testing.T) {
	p := &plateau.Plateau{Width: 5, Height: 5}
	r := NewRover(0, 0, South, p)
	err := r.Execute("M")
	if err == nil {
		t.Error("expected out of bounds error, got nil")
	}
}

func TestRover_Execute_RightTurnAndMove(t *testing.T) {
	p := &plateau.Plateau{Width: 5, Height: 5}
	r := NewRover(3, 3, East, p)
	err := r.Execute("MMRMMRMRRM")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.X != 5 || r.Y != 1 || r.D != East {
		t.Errorf("expected 5 1 E, got %d %d %s", r.X, r.Y, r.D)
	}
}
