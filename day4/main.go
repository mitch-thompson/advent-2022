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
	FILENAME   = "input.txt"
	CURRENTDIR = "."
)

type pair struct {
	firstElf, secondElf []int
}

func main() {
	filesys := os.DirFS(CURRENTDIR)
	f, err := filesys.Open(FILENAME)
	defer f.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	overlaps := sumOfOverlaps(f)
	fmt.Fprintln(os.Stdout, overlaps)
	f.Close()

	filesys = os.DirFS(CURRENTDIR)
	f, err = filesys.Open(FILENAME)
	defer f.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	anyOverlaps := sumOfAnyOverlaps(f)
	fmt.Fprintln(os.Stdout, anyOverlaps)
	f.Close()
}

func scanLines(fh fs.File) []pair {
	var pairs []pair

	fs := bufio.NewScanner(fh)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		line := fs.Text()
		if line != "" {
			var p pair
			firstElf := strings.Split(line, ",")[0]
			secondElf := strings.Split(line, ",")[1]
			p.firstElf = setElves(firstElf)
			p.secondElf = setElves(secondElf)
			pairs = append(pairs, p)
		}
	}
	return pairs
}

func sumOfOverlaps(f fs.File) int {
	pairs := scanLines(f)
	var sum int
	for _, p := range pairs {
		if p.isOverlap() {
			sum++
		}
	}
	return sum
}

func sumOfAnyOverlaps(f fs.File) int {
	pairs := scanLines(f)
	var sum int
	for _, p := range pairs {
		if p.isAnyOverlap() {
			sum++
		}
	}
	return sum
}

func setElves(elf string) []int {
	var elfCleaning []int
	i, err := strconv.Atoi(strings.Split(elf, "-")[0])
	j, err2 := strconv.Atoi(strings.Split(elf, "-")[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err2 != nil {
		fmt.Fprintln(os.Stderr, err2)
		os.Exit(1)
	}

	for i <= j {
		elfCleaning = append(elfCleaning, i)
		i++
	}

	return elfCleaning
}

func (p pair) isOverlap() bool {
	switch {
	case p.firstElf[0] >= p.secondElf[0] &&
		p.firstElf[len(p.firstElf)-1] <= p.secondElf[len(p.secondElf)-1]:
		return true
	case p.secondElf[0] >= p.firstElf[0] &&
		p.secondElf[len(p.secondElf)-1] <= p.firstElf[len(p.firstElf)-1]:
		return true
	}
	return false
}

func (p *pair) isAnyOverlap() bool {
	switch {
	case p.isOverlap():
		return true
	case p.firstElf[0] <= p.secondElf[len(p.secondElf)-1] &&
		p.firstElf[0] >= p.secondElf[0]:
		return true
	case p.firstElf[len(p.firstElf)-1] <= p.secondElf[len(p.secondElf)-1] &&
		p.firstElf[len(p.firstElf)-1] >= p.secondElf[0]:
		return true
	case p.secondElf[0] <= p.firstElf[len(p.firstElf)-1] &&
		p.secondElf[0] >= p.firstElf[0]:
		return true
	case p.secondElf[len(p.secondElf)-1] <= p.firstElf[len(p.firstElf)-1] &&
		p.secondElf[len(p.secondElf)-1] >= p.firstElf[0]:
		return true
	}
	return false
}
