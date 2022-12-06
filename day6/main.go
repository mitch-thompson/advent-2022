package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

const (
	FILENAME               = "input.txt"
	CURRENTDIR             = "."
	PACKET_WIDTH           = 4
	START_OF_MESSAGE_WIDTH = 14
)

type packet struct {
	input       string
	startMarker int
}

func main() {
	filesys := os.DirFS(CURRENTDIR)
	f, err := filesys.Open(FILENAME)
	defer f.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	p := readline(f)
	packetStart := p.startOfPacket(PACKET_WIDTH)
	messageStart := p.startOfPacket(START_OF_MESSAGE_WIDTH)
	fmt.Fprintln(os.Stdout, packetStart)
	fmt.Fprintln(os.Stdout, messageStart)
}

func readline(fh fs.File) packet {
	var p packet
	fs := bufio.NewScanner(fh)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		line := fs.Text()
		p.input = line
	}
	return p
}

func (p *packet) startOfPacket(width int) int {
	lineSlice := strings.Split(p.input, "")
	i := 0
	for i < len(lineSlice) {
		marker := lineSlice[i : width+i]
		finished := true
		for len(marker) > 0 {
			c := marker[0]
			marker = marker[1:]
			if contains(marker, c) {
				finished = false
				marker = nil
			}
		}

		if finished {
			return i + width
		}
		i++
	}
	return -1
}

func contains(ss []string, c string) bool {
	for _, s := range ss {
		if s == c {
			return true
		}
	}
	return false
}
