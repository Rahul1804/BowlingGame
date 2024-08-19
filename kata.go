package main

import (
	"errors"
	"fmt"
)

type Game struct {
	rolls        []int // Slice to store the rolls
	currentRoll  int   // Current roll number
	framesPlayed int   // Number of frames played (1-10)
	isGameOver   bool  // Flag to check if the game is over
}

// Initialize a new Game
func NewGame() *Game {
	return &Game{
		rolls: make([]int, 21), // Max rolls possible in a game (including bonus rolls in the 10th frame)
	}
}

// Roll method records the number of pins knocked down on a roll
func (g *Game) Roll(pins int) error {
	if g.isGameOver {
		return errors.New("the game is already over")
	}
	if pins < 0 || pins > 10 {
		return errors.New("invalid number of pins")
	}

	g.rolls[g.currentRoll] = pins
	g.currentRoll++

	// Frames 1-9 logic
	if g.framesPlayed < 9 {
		if pins == 10 { // Strike, move to the next frame immediately
			g.framesPlayed++
		} else if g.currentRoll%2 == 0 { // Two rolls complete the frame
			g.framesPlayed++
		}
	} else { // 10th frame logic
		if g.currentRoll > 18 { // Special handling for the 10th frame
			g.isGameOver = g.checkGameOverInFinalFrame()
		}
	}

	return nil
}

// checkGameOverInFinalFrame determines if the game is over after the 10th frame
func (g *Game) checkGameOverInFinalFrame() bool {
	if g.rolls[18] == 10 { // Strike in the first roll of the 10th frame
		return g.currentRoll >= 21
	} else if g.rolls[18]+g.rolls[19] == 10 { // Spare in the 10th frame
		return g.currentRoll >= 21
	} else { // No strike or spare in the 10th frame
		return g.currentRoll >= 20
	}
}

// Score method calculates the final score of the game
func (g *Game) Score() int {
	score := 0
	rollIndex := 0

	for frame := 0; frame < 10; frame++ {
		if g.isStrike(rollIndex) { // Strike logic
			score += 10 + g.rolls[rollIndex+1] + g.rolls[rollIndex+2]
			rollIndex++
		} else if g.isSpare(rollIndex) { // Spare logic
			score += 10 + g.rolls[rollIndex+2]
			rollIndex += 2
		} else { // Regular frame
			score += g.rolls[rollIndex] + g.rolls[rollIndex+1]
			rollIndex += 2
		}
	}

	return score
}

// isStrike checks if the roll is a strike
func (g *Game) isStrike(rollIndex int) bool {
	return g.rolls[rollIndex] == 10
}

// isSpare checks if the frame is a spare
func (g *Game) isSpare(rollIndex int) bool {
	return g.rolls[rollIndex]+g.rolls[rollIndex+1] == 10
}

// Main function to simulate a perfect game
func main() {
	game := NewGame()

	// Simulate a perfect game
	for i := 0; i < 12; i++ {
		err := game.Roll(10)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	fmt.Println("Final score:", game.Score()) // Should print 300 for a perfect game
}
