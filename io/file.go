package io

import (
  "bufio"
  "os"
  "strings"
)

type FileIo struct {
  path string
}

func (s FileIo) ReadLines() ([]string, error) {
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

func (s FileIo) WriteLines(lines []string) error {
  return s.doWriteLines(lines, os.O_TRUNC)
}

func (s FileIo) AppendLines(lines []string) error {
  return s.doWriteLines(lines, os.O_APPEND)
}

func (s FileIo) doWriteLines(lines []string, mode int) error {
  if len(lines) == 0 { return nil }
  file, err := os.OpenFile(s.path, os.O_WRONLY | os.O_CREATE | mode, 0644)
  defer file.Close()
  if err != nil { return err }
  file.Write([]byte(strings.Join(lines, "\n") + "\n"))
  return nil
}

