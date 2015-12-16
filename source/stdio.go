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
  if len(lines) == 0 { return nil }
  io := os.Stdout
  io.Write([]byte(strings.Join(lines, "\n") + "\n"))
  return nil
}

func (s StdioSource) AppendLines(lines []string) error {
  return s.WriteLines(lines)
}
