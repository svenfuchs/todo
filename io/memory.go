package io

import (
  "strings"
)

type MemoryIo struct {
  content string
}

func NewMemoryIo(content string) MemoryIo {
  return MemoryIo { content }
}

func (s *MemoryIo) ReadLines() ([]string, error) {
  return strings.Split(s.content, "\n"), nil
}

func (s *MemoryIo) WriteLines(lines []string) error {
  s.content = strings.Join(lines, "\n")
  return nil
}

func (s *MemoryIo) AppendLines(lines []string) error {
  s.content += strings.Join(lines, "\n")
  return nil
}
