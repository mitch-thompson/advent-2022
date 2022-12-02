package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

const (
	FILENAME   = "input.txt"
	CURRENTDIR = "."
	ROCK       = "Rock"
	PAPER      = "Paper"
	SCISSORS   = "Scissors"
	WIN        = "Win"
	LOSE       = "Lose"
	DRAW       = "Draw"
)

type game struct {
	yourChoice, opponentChoice string
	didWin                     string
	score                      int
}

func prs() []string {
	return []string{PAPER, ROCK, SCISSORS}
}

func prsKey(s string) int {
	for i, v := range prs() {
		if v == s {
			return i
		}
	}
	return -1
}

func main() {
	filesys := os.DirFS(CURRENTDIR)
	f, err := filesys.Open(FILENAME)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	score := sumOfGames(f)
	fmt.Fprintln(os.Stdout, score)
	f.Close()

	// Method 2
	f, err = filesys.Open(FILENAME)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	score = sumSecondColumnCorrected(f)
	fmt.Fprintln(os.Stdout, score)
}

func sumOfGames(f fs.File) int {
	games := *mapGames(f)
	for i, _ := range games {
		games[i].rockPaperOrScissors()
		games[i].cardPoints()
		games[i].playGame()
		games[i].gameScore()
	}

	return calculateScore(games)
}

func sumSecondColumnCorrected(f fs.File) int {
	games := *mapGamesAgain(f)
	for i, _ := range games {
		games[i].rockPaperOrScissors()
		games[i].gameScore()
		games[i].myChoose()
		games[i].cardPoints()
	}

	return calculateScore(games)
}

func calculateScore(games []game) int {
	var score int
	for _, g := range games {
		score += g.score
	}
	return score
}

func mapGamesAgain(fh fs.File) *[]game {
	var games []game

	fs := bufio.NewScanner(fh)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		var g game
		line := fs.Text()
		if line != "" {
			lineSlice := strings.Split(line, " ")
			g.opponentChoice = lineSlice[0]
			switch lineSlice[1] {
			case "X":
				g.didWin = LOSE
				break
			case "Y":
				g.didWin = DRAW
				break
			case "Z":
				g.didWin = WIN
				break
			}
			games = append(games, g)
		}
	}
	return &games
}

func mapGames(fh fs.File) *[]game {
	var games []game

	fs := bufio.NewScanner(fh)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		var g game
		line := fs.Text()
		if line != "" {
			lineSlice := strings.Split(line, " ")
			g.opponentChoice = lineSlice[0]
			g.yourChoice = lineSlice[1]
			games = append(games, g)
		}
	}

	return &games
}

func (g *game) rockPaperOrScissors() {
	switch g.yourChoice {
	case "X":
		g.yourChoice = ROCK
		break
	case "Y":
		g.yourChoice = PAPER
		break
	case "Z":
		g.yourChoice = SCISSORS
		break
	}
	switch g.opponentChoice {
	case "A":
		g.opponentChoice = ROCK
		break
	case "B":
		g.opponentChoice = PAPER
		break
	case "C":
		g.opponentChoice = SCISSORS
		break
	}
}

func (g *game) playGame() {
	switch {
	case g.yourChoice == g.opponentChoice:
		g.didWin = DRAW
		break
	case g.yourChoice == ROCK && g.opponentChoice == PAPER:
		g.didWin = LOSE
		break
	case g.yourChoice == ROCK && g.opponentChoice == SCISSORS:
		g.didWin = WIN
		break
	case g.yourChoice == SCISSORS && g.opponentChoice == ROCK:
		g.didWin = LOSE
		break
	case g.yourChoice == SCISSORS && g.opponentChoice == PAPER:
		g.didWin = WIN
		break
	case g.yourChoice == PAPER && g.opponentChoice == SCISSORS:
		g.didWin = LOSE
		break
	case g.yourChoice == PAPER && g.opponentChoice == ROCK:
		g.didWin = WIN
		break
	}
}

func (g *game) cardPoints() {
	switch g.yourChoice {
	case PAPER:
		g.score += 2
		break
	case ROCK:
		g.score += 1
		break
	case SCISSORS:
		g.score += 3
		break
	}
}

func (g *game) gameScore() {
	switch g.didWin {
	case WIN:
		g.score += 6
		break
	case LOSE:
		g.score += 0
	case DRAW:
		g.score += 3
	}
}

func (g *game) myChoose() {
	switch {
	case g.didWin == DRAW:
		g.yourChoice = g.opponentChoice
		break
	case g.didWin == LOSE:
		i := prsKey(g.opponentChoice)
		if i >= len(prs())-1 {
			g.yourChoice = prs()[0]
		} else {
			g.yourChoice = prs()[i+1]
		}
		break
	case g.didWin == WIN:
		i := prsKey(g.opponentChoice)
		if i == 0 {
			g.yourChoice = prs()[len(prs())-1]
		} else {
			g.yourChoice = prs()[i-1]
		}
		break
	}
}
