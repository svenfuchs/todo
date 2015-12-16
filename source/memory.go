package source

import (
  "strings"
)

type MemorySource struct {
  content string
}

func NewMemorySource(content string) MemorySource {
  return MemorySource { content }
}

func (s *MemorySource) ReadLines() ([]string, error) {
  return strings.Split(s.content, "\n"), nil
}

func (s *MemorySource) WriteLines(lines []string) error {
  s.content = strings.Join(lines, "\n")
  return nil
}

func (s *MemorySource) AppendLines(lines []string) error {
  s.content += strings.Join(lines, "\n")
  return nil
}
