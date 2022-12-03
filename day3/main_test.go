package main

import (
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

//Each rune is an item
//each rucksack is a line
//half way part marks second compartment
//a-z 1-26, A-Z 27-52

const testInput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func testSetup() fs.File {
	fs := fstest.MapFS{
		"input.txt": {Data: []byte(testInput)},
	}

	f, _ := fs.Open("input.txt")
	return f
}

func TestIntegrationPrioritySum(t *testing.T) {
	f := testSetup()
	defer f.Close()

	got := sumOfPriorities(f)
	want := 157

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestIntegrationGropuSum(t *testing.T) {
	f := testSetup()
	defer f.Close()

	got := sumOfGroups(f)
	want := 70

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestWhatIsInCompartments(t *testing.T) {
	cases := []struct {
		testName  string
		pi        packedItems
		wantLeft  []rune
		wantRight []rune
	}{
		{
			"Even",
			packedItems{
				allItems: []rune{'a', 'b', 'c', 'd'},
			},
			[]rune{'a', 'b'},
			[]rune{'c', 'd'},
		},
	}
	for _, c := range cases {
		t.Run("test "+c.testName, func(t *testing.T) {
			c.pi.whatIsInCompartments()
			if !reflect.DeepEqual(c.pi.compartmentOne, c.wantLeft) {
				t.Errorf("Got %v, want %v", c.pi.compartmentOne, c.wantLeft)
			}

			if !reflect.DeepEqual(c.pi.compartmentTwo, c.wantRight) {
				t.Errorf("Got %v, want %v", c.pi.compartmentTwo, c.wantRight)
			}
		})
	}
}

func TestGetDuplicates(t *testing.T) {
	cases := []struct {
		testName string
		pi       packedItems
		want     []rune
	}{
		{
			"ZeroDupliate",
			packedItems{
				compartmentOne: []rune{'a', 'b', 'c'},
				compartmentTwo: []rune{'g', 'r', 'f'},
			},
			nil,
		},
		{
			"OneDupliate",
			packedItems{
				compartmentOne: []rune{'a', 'b', 'c'},
				compartmentTwo: []rune{'g', 'b', 'f'},
			},
			[]rune{'b'},
		},
		{
			"TwoDupliate",
			packedItems{
				compartmentOne: []rune{'a', 'b', 'c'},
				compartmentTwo: []rune{'g', 'b', 'a'},
			},
			[]rune{'a', 'b'},
		},
	}
	for _, c := range cases {
		t.Run("test "+c.testName, func(t *testing.T) {
			c.pi.getDuplicates()
			if !reflect.DeepEqual(c.pi.duplicates, c.want) {
				t.Errorf("%v failed got %v wanted %v", c.testName, c.pi.duplicates, c.want)
			}
		})
	}
}

func TestGetPrioritySum(t *testing.T) {
	pi := packedItems{duplicates: []rune{'a', 'b', 'c', 'd'}}
	want := 10

	pi.getPrioritySum()
	if pi.prioritySum != want {
		t.Errorf("Expected %v, got %v", want, pi.prioritySum)
	}
}
