package io

import (
  "os"
)

type Io interface {
  ReadLines() []string
  WriteLines([]string)
  AppendLines([]string)
}

func NewIo(path string) Io {
  stat, _ := os.Stdin.Stat()
  if (stat.Mode() & os.ModeCharDevice) == 0 {
    return NewStdOut()
  } else {
    return NewFileIo(path)
  }
}
