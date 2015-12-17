package io

import (
  "bufio"
  "os"
  "strings"
)

// var stdin *os.File

func NewStdOut() Io {
  return StdIo{}
}

func NewStdErr() Io {
  return StdIo{ os.Stderr }
}

type StdIo struct{
  out *os.File
}

func (s StdIo) ReadLines() ([]string, error) {
  lines   := []string{}
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func (s StdIo) WriteLines(lines []string) error {
  if len(lines) == 0 { return nil }
  s.out.Write([]byte(strings.Join(lines, "\n") + "\n"))
  return nil
}

func (s StdIo) AppendLines(lines []string) error {
  return s.WriteLines(lines)
}
