package io

import (
  "strings"
)

func NewMemoryIo(content string) Io {
  return MemoryIo { &content }
}

type MemoryIo struct {
  content *string
}

func (s MemoryIo) ReadLines() ([]string, error) {
  return strings.Split(*s.content, "\n"), nil
}

func (s MemoryIo) WriteLines(lines []string) error {
  *s.content = strings.Join(lines, "\n")
  return nil
}

func (s MemoryIo) AppendLines(lines []string) error {
  *s.content = *s.content + strings.Join(lines, "\n")
  return nil
}
