package todo

import (
  "bufio"
  "os"
  "strings"
)

// var stdin *os.File

type Source interface {
  ReadLines() ([]string, error)
  MustReadLines() ([]string)
  WriteLines([]string) error
  MustWriteLines([]string)
}

type FileSource struct {
  path string
}

func NewFileSource(path string) Source {
  return FileSource { path: path }
}

func (f FileSource) ReadLines() ([]string, error) {
  var (
    err error
    io *os.File
  )

  if f.path == "" {
    io = os.Stdin
  } else {
    io, err = os.Open(f.path)
    defer io.Close()
  }

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

func (f FileSource) MustReadLines() []string {
  lines, err := f.ReadLines()
  check(err)
  return lines
}

func (f FileSource) WriteLines(lines []string) error {
  var (
    err error
    io *os.File
  )

  if f.path == "" {
    io = os.Stdout
  } else {
    io, err = os.OpenFile(f.path, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
    defer io.Close()
  }

  if err == nil {
    io.Write([]byte(strings.Join(lines, "\n") + "\n"))
  }
  return err
}

func (f FileSource) MustWriteLines(lines []string) {
  err := f.WriteLines(lines)
  check(err)
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}

type MemorySource struct {
  content string
}

func NewMemorySource(content string) MemorySource {
  return MemorySource { content }
}

func (s *MemorySource) ReadLines() ([]string, error) {
  return strings.Split(s.content, "\n"), nil
}

func (s *MemorySource) MustReadLines() []string {
  lines, _ := s.ReadLines()
  return lines
}

func (s *MemorySource) WriteLines(lines []string) error {
  s.content = strings.Join(lines, "\n")
  return nil
}

func (s *MemorySource) MustWriteLines(lines []string) {
  s.WriteLines(lines)
}
