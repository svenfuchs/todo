package source

import (
  "bufio"
  "os"
  "strings"
)

// var stdin *os.File

type StdioSource struct{}

func (s StdioSource) ReadLines() ([]string, error) {
  lines   := []string{}
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func (s StdioSource) WriteLines(lines []string) error {
  io := os.Stdout
  io.Write([]byte(strings.Join(lines, "\n") + "\n"))
  return nil
}

func (s StdioSource) MustReadLines() []string {
  return mustReadLines(s)
}

func (s StdioSource) MustWriteLines(lines []string) {
  mustWriteLines(s, lines)
}

