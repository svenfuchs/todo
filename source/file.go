package source

import (
  "bufio"
  "os"
  "strings"
)

type FileSource struct {
  path string
}

func (s FileSource) ReadLines() ([]string, error) {
  io, err := os.Open(s.path)
  defer io.Close()

  lines   := []string{}
  scanner := bufio.NewScanner(io)

  if err != nil {
    return lines, err
  }

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func (s FileSource) WriteLines(lines []string) error {
  io, err := os.OpenFile(s.path, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
  if err == nil { return err }
  io.Write([]byte(strings.Join(lines, "\n") + "\n"))
  return nil
}

func (s FileSource) MustReadLines() []string {
  return mustReadLines(s)
}

func (s FileSource) MustWriteLines(lines []string) {
  mustWriteLines(s, lines)
}

