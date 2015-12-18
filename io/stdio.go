package io

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

var (
	Stdin  io.ReadWriter = os.Stdin
	Stdout io.ReadWriter = os.Stdout
	Stderr io.ReadWriter = os.Stderr
)

// type file interface {
//   io.Closer
//   io.Reader
//   io.Writer
// }

func NewStdOut() Io {
	return StdIo{Stdout}
}

func NewStdErr() Io {
	return StdIo{Stderr}
}

type StdIo struct {
	out io.Writer
}

func (s StdIo) ReadLines() []string {
	lines := []string{}
	scanner := bufio.NewScanner(Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func (s StdIo) WriteLines(lines []string) {
	if len(lines) == 0 {
		return
	}
	str := strings.Join(lines, "\n") + "\n"
	s.out.Write([]byte(str))
}

func (s StdIo) AppendLines(lines []string) {
	s.WriteLines(lines)
}
