package io

import (
  "bufio"
  "log"
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

func (s StdIo) ReadLines() []string {
  lines   := []string{}
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return lines
}

func (s StdIo) WriteLines(lines []string) {
  if len(lines) == 0 { return }
  str := strings.Join(lines, "\n") + "\n"
  s.out.Write([]byte(str))
}

func (s StdIo) AppendLines(lines []string) {
  s.WriteLines(lines)
}
