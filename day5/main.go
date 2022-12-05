package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

const (
	FILENAME         = "input.txt"
	CURRENTDIR       = "."
	MOVE             = "move "
	COLUMNS          = "["
	FROM             = "from"
	TO               = "to"
	RUNES_PER_COLUMN = 4
)

type crates struct {
	crateSlices                  [9][]string
	moveTo, moveAmount, moveFrom []int
}

func main() {
	filesys := os.DirFS(CURRENTDIR)
	f, err := filesys.Open(FILENAME)
	defer f.Close()

	errHandler(err)

	top := topCrates(f)
	fmt.Fprintln(os.Stdout, top)
	f.Close()

	filesys = os.DirFS(CURRENTDIR)
	f, err = filesys.Open(FILENAME)
	defer f.Close()

	errHandler(err)

	topMultipleMove := topMultipleCrates(f)
	fmt.Fprintln(os.Stdout, topMultipleMove)
}

func errHandler(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func topCrates(f fs.File) string {
	var instructions string
	crateInstructions := scanLines(f)

	for i, _ := range crateInstructions.moveAmount {
		crateInstructions.moveCrate(i)
	}

	for _, s := range crateInstructions.crateSlices {
		if len(s) > 0 {
			instructions = instructions + s[0]
		}
	}

	return instructions
}

func topMultipleCrates(f fs.File) string {
	var instructions string
	crateInstructions := scanLines(f)

	for i, _ := range crateInstructions.moveAmount {
		crateInstructions.moveMultipleCrates(i)
	}

	for _, s := range crateInstructions.crateSlices {
		if len(s) > 0 {
			instructions = instructions + s[0]
		}
	}

	return instructions
}

func scanLines(fh fs.File) crates {
	var crateInstructions crates

	fs := bufio.NewScanner(fh)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		line := fs.Text()
		switch {
		case strings.Contains(line, MOVE):
			cutMove := strings.Split(line, MOVE)[1]
			amount, err := strconv.Atoi(strings.TrimSpace(strings.Split(cutMove, FROM)[0]))
			errHandler(err)
			crateInstructions.moveAmount = append(crateInstructions.moveAmount, amount)
			cutFrom := strings.Split(cutMove, FROM)[1]
			from, err := strconv.Atoi(strings.TrimSpace(strings.Split(cutFrom, TO)[0]))
			errHandler(err)
			to, err := strconv.Atoi(strings.TrimSpace(strings.Split(cutFrom, TO)[1]))
			errHandler(err)
			crateInstructions.moveFrom = append(crateInstructions.moveFrom, from)
			crateInstructions.moveTo = append(crateInstructions.moveTo, to)
			break
		case strings.Contains(line, COLUMNS):
			lineSlice := strings.Split(line, "")
			i := 0
			for i < len(lineSlice) {
				if lineSlice[i] == COLUMNS {
					crateInstructions.crateSlices[i/RUNES_PER_COLUMN] = append(crateInstructions.crateSlices[i/RUNES_PER_COLUMN], lineSlice[i+1])
				}
				i += RUNES_PER_COLUMN
			}
			break
		}
	}

	return crateInstructions
}

func (crateInstructions *crates) moveCrate(pos int) {
	moveTo := crateInstructions.moveTo[pos] - 1
	moveFrom := crateInstructions.moveFrom[pos] - 1
	amount := crateInstructions.moveAmount[pos]

	for i := 0; i < amount; i++ {
		crateInstructions.crateSlices[moveTo] = append([]string{crateInstructions.crateSlices[moveFrom][0]},
			crateInstructions.crateSlices[moveTo]...)
		crateInstructions.crateSlices[moveFrom] = crateInstructions.crateSlices[moveFrom][1:]
	}
}

func (crateInstructions *crates) moveMultipleCrates(pos int) {
	moveTo := crateInstructions.moveTo[pos] - 1
	moveFrom := crateInstructions.moveFrom[pos] - 1
	amount := crateInstructions.moveAmount[pos]

	var newCrates []string
	for i := 0; i < amount; i++ {
		newCrates = append(newCrates, crateInstructions.crateSlices[moveFrom][i])
	}
	crateInstructions.crateSlices[moveTo] = append(newCrates,
		crateInstructions.crateSlices[moveTo]...)

	if len(crateInstructions.crateSlices[moveFrom]) >= amount {
		crateInstructions.crateSlices[moveFrom] = crateInstructions.crateSlices[moveFrom][amount:]
	} else {
		crateInstructions.crateSlices[moveFrom] = []string{}
	}
}
