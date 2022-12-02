package main

import (
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

// Other
// A -> Rock
// B -> Paper
// C -> Scissors

// You
// X -> Rock
// Y -> Paper
// Z -> Scissors

// Rock 1
// Paper 2
// Scissors 3
// Loss 0
// Draw 3
// Win 6

const testInput = `
A Y
B X
C Z
`

func testSetup() fs.File {
	fs := fstest.MapFS{
		"input.txt": {Data: []byte(testInput)},
	}

	f, _ := fs.Open("input.txt")
	return f
}

func TestSumOfGames(t *testing.T) {
	f := testSetup()
	defer f.Close()

	got := sumOfGames(f)
	want := 15

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestSumOfSecondGames(t *testing.T) {
	f := testSetup()
	defer f.Close()

	got := sumSecondColumnCorrected(f)
	want := 12

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestMapGames(t *testing.T) {
	f := testSetup()
	defer f.Close()

	var g1, g2, g3 game
	g1.opponentChoice = "A"
	g1.yourChoice = "Y"
	g2.opponentChoice = "B"
	g2.yourChoice = "X"
	g3.opponentChoice = "C"
	g3.yourChoice = "Z"
	var want []game

	want = append(want, g1, g2, g3)
	got := mapGames(f)

	if !reflect.DeepEqual(want, *got) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestRockPaperOrScissors(t *testing.T) {
	var g1, g2, g3, g4, g5, g6 game

	g1.opponentChoice = "A"
	g1.yourChoice = "Y"
	g2.opponentChoice = "B"
	g2.yourChoice = "X"
	g3.opponentChoice = "C"
	g3.yourChoice = "Z"

	var want, got []game
	got = append(got, g1, g2, g3)

	g4.opponentChoice = "Rock"
	g4.yourChoice = "Paper"
	g5.opponentChoice = "Paper"
	g5.yourChoice = "Rock"
	g6.opponentChoice = "Scissors"
	g6.yourChoice = "Scissors"

	want = append(want, g4, g5, g6)

	for i, _ := range got {
		got[i].rockPaperOrScissors()
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestPlayGame(t *testing.T) {
	cases := []struct {
		testName  string
		g         game
		want      string
		wantScore int
	}{
		{
			"draw",
			game{
				yourChoice:     ROCK,
				opponentChoice: ROCK,
				didWin:         "",
				score:          0,
			},
			DRAW,
			3,
		},
		{
			"Rock vs Paper",
			game{
				yourChoice:     ROCK,
				opponentChoice: PAPER,
				didWin:         "",
				score:          0,
			},
			LOSE,
			0,
		},
		{
			"Rock vs Scissors",
			game{
				yourChoice:     ROCK,
				opponentChoice: SCISSORS,
				didWin:         "",
				score:          0,
			},
			WIN,
			6,
		},
		{
			"Scissors vs Rock",
			game{
				yourChoice:     SCISSORS,
				opponentChoice: ROCK,
				didWin:         "",
				score:          0,
			},
			LOSE,
			0,
		},
		{
			"Scissors vs Paper",
			game{
				yourChoice:     SCISSORS,
				opponentChoice: PAPER,
				didWin:         "",
				score:          0,
			},
			WIN,
			6,
		},
		{
			"Paper vs Scissors",
			game{
				yourChoice:     PAPER,
				opponentChoice: SCISSORS,
				didWin:         "",
				score:          0,
			},
			LOSE,
			0,
		},
		{
			"Paper vs Rock",
			game{
				yourChoice:     PAPER,
				opponentChoice: ROCK,
				didWin:         "",
				score:          0,
			},
			WIN,
			6,
		},
	}

	for _, c := range cases {
		t.Run("test "+c.testName, func(t *testing.T) {
			c.g.playGame()
			if c.g.didWin != c.want {
				t.Errorf("Got %v, want %v", c.g.score, c.want)
			}
		})
	}
}

func TestMapGamesAgain(t *testing.T) {

}

func TestMyChoose(t *testing.T) {

}

func TestCardPoints(t *testing.T) {

}

func TestGameScore(t *testing.T) {

}
