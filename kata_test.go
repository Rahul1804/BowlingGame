package main

import "testing"

// helper functions for testing
func rollMany(g *Game, n int, pins int) {
	for i := 0; i < n; i++ {
		_ = g.Roll(pins)
	}
}

func rollSpare(g *Game) {
	_ = g.Roll(5)
	_ = g.Roll(5)
}

func rollStrike(g *Game) {
	_ = g.Roll(10)
}

// unit tests
func TestGutterGame(t *testing.T) {
	game := NewGame()
	rollMany(game, 20, 0) // roll 20 times with 0 pins
	if game.Score() != 0 {
		t.Errorf("Expected score to be 0, but got %d", game.Score())
	}
}

func TestAllOnes(t *testing.T) {
	game := NewGame()
	rollMany(game, 20, 1) // roll 20 times with 1 pin
	if game.Score() != 20 {
		t.Errorf("Expected score to be 20, but got %d", game.Score())
	}
}

func TestOneSpare(t *testing.T) {
	game := NewGame()
	rollSpare(game)       // Roll a spare
	_ = game.Roll(3)      // Next roll
	rollMany(game, 17, 0) // Complete the game with 0s

	if game.Score() != 16 {
		t.Errorf("Expected score to be 16, but got %d", game.Score())
	}
}

func TestOneStrike(t *testing.T) {
	game := NewGame()
	rollStrike(game)      // Strike
	_ = game.Roll(3)      // Next roll
	_ = game.Roll(4)      // Next roll
	rollMany(game, 16, 0) // Complete the game with 0s

	if game.Score() != 24 {
		t.Errorf("Expected score to be 24, but got %d", game.Score())
	}
}

func TestPerfectGame(t *testing.T) {
	game := NewGame()
	rollMany(game, 12, 10) // Roll 12 strikes

	if game.Score() != 300 {
		t.Errorf("Expected score to be 300, but got %d", game.Score())
	}
}

func TestGameEndsProperly(t *testing.T) {
	game := NewGame()

	// Roll normal frames
	for i := 0; i < 18; i++ {
		_ = game.Roll(4)
	}

	_ = game.Roll(4)
	_ = game.Roll(5) // Final frame without bonuses
	if !game.isGameOver {
		t.Errorf("Expected game to be over, but it is not.")
	}

	// Test game ending with a spare in the 10th frame
	game = NewGame()
	for i := 0; i < 18; i++ {
		_ = game.Roll(4)
	}
	_ = game.Roll(5) // Final frame first roll
	_ = game.Roll(5) // Spare
	_ = game.Roll(4) // Bonus roll
	if !game.isGameOver {
		t.Errorf("Expected game to be over after bonus roll, but it is not.")
	}
}
