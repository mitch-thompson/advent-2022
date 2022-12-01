package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strconv"
)

const (
	FILENAME   = "input.txt"
	CURRENTDIR = "."
)

func main() {
	filesys := os.DirFS(CURRENTDIR)
	f, err := filesys.Open(FILENAME)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}
	highestCalories := highest(f)
	fmt.Fprintln(os.Stdout, highestCalories)
	f.Close()

	f, err = filesys.Open(FILENAME)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}
	highestThreeCalories := highestThree(f)
	fmt.Fprintln(os.Stdout, highestThreeCalories)
}

/**
highest returns calorie count of highest elf
*/
func highest(fh fs.File) int {
	elves := scanLines(fh)
	highestCalories := 0

	for _, calories := range elves {
		if countCalories(calories) > highestCalories {
			highestCalories = countCalories(calories)
		}
	}

	return highestCalories
}

/**
highestThree returns sum of the highest three elves calorie count
*/
func highestThree(fh fs.File) int {
	elves := scanLines(fh)
	var elfTotalCalories []int

	for _, elf := range elves {
		elfTotalCalories = append(elfTotalCalories, countCalories(elf))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfTotalCalories)))

	sum := 0
	for j := 0; j < 3; j++ {
		sum += elfTotalCalories[j]
	}
	return sum
}

/**
scanLines scans file line by line returning slice of slice of ints
using an empty line as a deliminator
*/
func scanLines(fh fs.File) [][]int {
	var elves [][]int
	var calories []int

	fs := bufio.NewScanner(fh)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		line := fs.Text()
		if line == "" {
			elves = append(elves, calories)
			calories = nil
		} else {
			lineInt, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}
			calories = append(calories, lineInt)
		}
	}
	return elves
}

/**
countCalories returns sum of calories from an int slice
*/
func countCalories(is []int) int {
	i := 0
	for _, j := range is {
		i += j
	}
	return i
}
