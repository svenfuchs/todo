package source

import (
  "bufio"
  "os"
  "strings"
)

// var stdin *os.File

func New(path string) Io {
  return File { path: path }
}

type Io interface {
  ReadLines() ([]string, error)
  MustReadLines() ([]string)
  WriteLines([]string) error
  MustWriteLines([]string)
}

type File struct {
  path string
}

func (f File) ReadLines() ([]string, error) {
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

func (f File) MustReadLines() []string {
  lines, err := f.ReadLines()
  check(err)
  return lines
}

func (f File) WriteLines(lines []string) error {
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

func (f File) MustWriteLines(lines []string) {
  err := f.WriteLines(lines)
  check(err)
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}

type Memory struct {
  Content string
}

func (s *Memory) ReadLines() ([]string, error) {
  return strings.Split(s.Content, "\n"), nil
}

func (s *Memory) MustReadLines() []string {
  lines, _ := s.ReadLines()
  return lines
}

func (s *Memory) WriteLines(lines []string) error {
  s.Content = strings.Join(lines, "\n")
  return nil
}

func (s *Memory) MustWriteLines(lines []string) {
  s.WriteLines(lines)
}
