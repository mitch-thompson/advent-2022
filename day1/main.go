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

func highestThree(fh fs.File) int {
	elves := scanLines(fh)
	var elfTotalCalories []int

	for _, elf := range elves {
		var sum int
		for _, value := range elf {
			sum += value
		}
		elfTotalCalories = append(elfTotalCalories, sum)
	}

	sort.Ints(elfTotalCalories)
	sort.Sort(sort.Reverse(sort.IntSlice(elfTotalCalories)))

	sum := 0
	for j := 0; j < 3; j++ {
		sum += elfTotalCalories[j]
	}
	return sum
}

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

func countCalories(is []int) int {
	i := 0
	for _, j := range is {
		i += j
	}
	return i
}
