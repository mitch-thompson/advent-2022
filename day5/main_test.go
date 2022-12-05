package main

import (
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

const testInput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func testSetup() fs.File {
	fs := fstest.MapFS{
		"input.txt": {Data: []byte(testInput)},
	}

	f, _ := fs.Open("input.txt")
	return f
}

func TestIntegrationTopCrates(t *testing.T) {
	f := testSetup()
	defer f.Close()

	got := topCrates(f)
	want := "CMZ"

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestIntegrationMultipleCrates(t *testing.T) {
	f := testSetup()
	defer f.Close()

	got := topMultipleCrates(f)
	want := "MCD"

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestScanLines(t *testing.T) {
	f := testSetup()
	defer f.Close()

	got := scanLines(f)
	want := crates{
		crateSlices: [9][]string{[]string{"N", "Z"}, []string{"D", "C", "M"}, []string{"P"}},
		moveTo:      []int{1, 3, 1, 2},
		moveAmount:  []int{1, 3, 2, 1},
		moveFrom:    []int{2, 1, 2, 1},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestMoveCrates(t *testing.T) {
	got := crates{
		crateSlices: [9][]string{[]string{"N", "Z"}, []string{"D", "C", "M"}, []string{"P"}},
		moveTo:      []int{1, 3, 1, 2},
		moveAmount:  []int{1, 3, 2, 1},
		moveFrom:    []int{2, 1, 2, 1},
	}

	pos := 0

	want := crates{
		crateSlices: [9][]string{[]string{"D", "N", "Z"}, []string{"C", "M"}, []string{"P"}},
		moveTo:      []int{1, 3, 1, 2},
		moveAmount:  []int{1, 3, 2, 1},
		moveFrom:    []int{2, 1, 2, 1},
	}

	got.moveCrate(pos)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestMoveMultipleCrates(t *testing.T) {
	cases := []struct {
		testName string
		pos      int
		want     crates
		got      crates
	}{
		{
			"First Position",
			0,
			crates{
				crateSlices: [9][]string{[]string{"D", "N", "Z"}, []string{"C", "M"}, []string{"P"}},
				moveTo:      []int{1, 3, 1, 2},
				moveAmount:  []int{1, 3, 2, 1},
				moveFrom:    []int{2, 1, 2, 1},
			},
			crates{
				crateSlices: [9][]string{[]string{"N", "Z"}, []string{"D", "C", "M"}, []string{"P"}},
				moveTo:      []int{1, 3, 1, 2},
				moveAmount:  []int{1, 3, 2, 1},
				moveFrom:    []int{2, 1, 2, 1},
			},
		},
		{
			"Second Position",
			1,
			crates{
				crateSlices: [9][]string{[]string{}, []string{"C", "M"}, []string{"D", "N", "Z", "P"}},
				moveTo:      []int{1, 3, 1, 2},
				moveAmount:  []int{1, 3, 2, 1},
				moveFrom:    []int{2, 1, 2, 1},
			},
			crates{
				crateSlices: [9][]string{[]string{"D", "N", "Z"}, []string{"C", "M"}, []string{"P"}},
				moveTo:      []int{1, 3, 1, 2},
				moveAmount:  []int{1, 3, 2, 1},
				moveFrom:    []int{2, 1, 2, 1},
			},
		},
		{
			"Third Position",
			2,
			crates{
				crateSlices: [9][]string{[]string{"C", "M"}, []string{}, []string{"D", "N", "Z", "P"}},
				moveTo:      []int{1, 3, 1, 2},
				moveAmount:  []int{1, 3, 2, 1},
				moveFrom:    []int{2, 1, 2, 1},
			},
			crates{
				crateSlices: [9][]string{[]string{}, []string{"C", "M"}, []string{"D", "N", "Z", "P"}},
				moveTo:      []int{1, 3, 1, 2},
				moveAmount:  []int{1, 3, 2, 1},
				moveFrom:    []int{2, 1, 2, 1},
			},
		},
		{
			"Fourth Position",
			3,
			crates{
				crateSlices: [9][]string{[]string{"M"}, []string{"C"}, []string{"D", "N", "Z", "P"}},
				moveTo:      []int{1, 3, 1, 2},
				moveAmount:  []int{1, 3, 2, 1},
				moveFrom:    []int{2, 1, 2, 1},
			},
			crates{
				crateSlices: [9][]string{[]string{"C", "M"}, []string{}, []string{"D", "N", "Z", "P"}},
				moveTo:      []int{1, 3, 1, 2},
				moveAmount:  []int{1, 3, 2, 1},
				moveFrom:    []int{2, 1, 2, 1},
			},
		},
	}
	for _, c := range cases {
		t.Run("test "+c.testName, func(t *testing.T) {
			c.got.moveMultipleCrates(c.pos)
			if !reflect.DeepEqual(c.got, c.want) {
				t.Errorf("Expected %v, got %v", c.want, c.got)
			}
		})
	}
}
