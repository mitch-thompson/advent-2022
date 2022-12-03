package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

const (
	FILENAME   = "input.txt"
	CURRENTDIR = "."
)

type packedItems struct {
	allItems, compartmentOne, compartmentTwo, duplicates []rune
	prioritySum                                          int
}

func valueMap() map[rune]int {
	return map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'j': 10,
		'k': 11,
		'l': 12,
		'm': 13,
		'n': 14,
		'o': 15,
		'p': 16,
		'q': 17,
		'r': 18,
		's': 19,
		't': 20,
		'u': 21,
		'v': 22,
		'w': 23,
		'x': 24,
		'y': 25,
		'z': 26,
		'A': 27,
		'B': 28,
		'C': 29,
		'D': 30,
		'E': 31,
		'F': 32,
		'G': 33,
		'H': 34,
		'I': 35,
		'J': 36,
		'K': 37,
		'L': 38,
		'M': 39,
		'N': 40,
		'O': 41,
		'P': 42,
		'Q': 43,
		'R': 44,
		'S': 45,
		'T': 46,
		'U': 47,
		'V': 48,
		'W': 49,
		'X': 50,
		'Y': 51,
		'Z': 52,
	}
}

func main() {
	filesys := os.DirFS(CURRENTDIR)
	f, err := filesys.Open(FILENAME)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	sumPriorities := sumOfPriorities(f)
	fmt.Fprintln(os.Stdout, sumPriorities)
	f.Close()

	filesys = os.DirFS(CURRENTDIR)
	f, err = filesys.Open(FILENAME)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	sumGroups := sumOfGroups(f)
	fmt.Fprintln(os.Stdout, sumGroups)
}

func sumOfPriorities(f fs.File) int {
	pi := scanLines(f)
	var sum int
	for i, _ := range pi {
		pi[i].whatIsInCompartments()
		pi[i].getDuplicates()
		pi[i].getPrioritySum()
		sum += pi[i].prioritySum
	}
	return sum
}

func sumOfGroups(f fs.File) int {
	pi := scanLines(f)
	var sum int
	for i := 0; i < len(pi); {
		packOne := pi[i]
		packTwo := pi[i+1]
		packThree := pi[i+2]

		r, err := checkDuplicatesGroup(packOne, packTwo, packThree)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		sum += valueMap()[r]
		i += 3
	}
	return sum
}

func checkDuplicatesGroup(packOne, packTwo, packThree packedItems) (rune, error) {
	for _, r1 := range packOne.allItems {
		for _, r2 := range packTwo.allItems {
			for _, r3 := range packThree.allItems {
				if r1 == r2 && r2 == r3 {
					return r1, nil
				}
			}
		}
	}
	return '0', errors.New("unmatched Rune")
}

func scanLines(fh fs.File) []packedItems {
	var packed []packedItems

	fs := bufio.NewScanner(fh)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		line := fs.Text()
		if line != "" {
			var p packedItems
			p.allItems = []rune(line)
			packed = append(packed, p)
		}
	}
	return packed
}

func (pi *packedItems) whatIsInCompartments() {
	for i, r := range pi.allItems {
		if i < len(pi.allItems)/2 {
			pi.compartmentOne = append(pi.compartmentOne, r)
		} else {
			pi.compartmentTwo = append(pi.compartmentTwo, r)
		}
	}
}

func (pi *packedItems) getDuplicates() {
	for _, r := range pi.compartmentOne {
		for _, s := range pi.compartmentTwo {
			if r == s && !pi.isPresent(r) {
				pi.duplicates = append(pi.duplicates, r)
			}
		}
	}
}

func (pi *packedItems) getPrioritySum() {
	for _, r := range pi.duplicates {
		pi.prioritySum += valueMap()[r]
	}
}

func (pi *packedItems) isPresent(r rune) bool {
	for _, s := range pi.duplicates {
		if s == r {
			return true
		}
	}
	return false
}
