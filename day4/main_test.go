package main

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

const testInput = `
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func testSetup() fs.File {
	fs := fstest.MapFS{
		"input.txt": {Data: []byte(testInput)},
	}

	f, _ := fs.Open("input.txt")
	return f
}

func TestIntegrationFirstProblem(t *testing.T) {
	f := testSetup()
	defer f.Close()

	got := sumOfOverlaps(f)
	want := 2

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestIntegrationSecondProblem(t *testing.T) {
	f := testSetup()
	defer f.Close()

	got := sumOfAnyOverlaps(f)
	want := 4

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
