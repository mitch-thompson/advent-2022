package main

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

const testInput = ``

func testSetup() fs.File {
	fs := fstest.MapFS{
		"input.txt": {Data: []byte(testInput)},
	}

	f, _ := fs.Open("input.txt")
	return f
}

func TestIntegrationTestCharactersToStart(t *testing.T) {
	cases := []struct {
		testName     string
		input        string
		packetStart  int
		messageStart int
	}{
		{
			"First Test",
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			7,
			19,
		},
		{
			"Second Test",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			5,
			23,
		},
		{
			"Third Test",
			"nppdvjthqldpwncqszvftbrmjlhg",
			6,
			23,
		},
		{
			"Fourth Test",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			10,
			29,
		},
		{
			"Fifth Test",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			11,
			26,
		},
	}

	for _, c := range cases {
		p := packet{input: c.input}

		got := p.startOfPacket(PACKET_WIDTH)
		if got != c.packetStart {
			t.Errorf("Expected %v, got %v", c.packetStart, got)
		}

		got = p.startOfPacket(START_OF_MESSAGE_WIDTH)
		if got != c.messageStart {
			t.Errorf("Expected %v, got %v", c.messageStart, got)
		}
	}
}
