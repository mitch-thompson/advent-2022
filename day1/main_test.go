package main

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

const testInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

`

func testSetup() fs.File {
	fs := fstest.MapFS{
		"input.txt": {Data: []byte(testInput)},
	}

	f, _ := fs.Open("input.txt")
	return f
}

func TestHighest(t *testing.T) {
	f := testSetup()
	defer f.Close()

	got := highest(f)
	want := 24000

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestHighestThree(t *testing.T) {
	f := testSetup()
	defer f.Close()

	want := 24000 + 11000 + 10000

	got := highestThree(f)

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
