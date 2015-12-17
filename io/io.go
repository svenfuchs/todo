package io

// var stdin *os.File

import (
  "os"
)

type Io interface {
  ReadLines() ([]string, error)
  WriteLines([]string) error
  AppendLines([]string) error
}

func NewIo(path string) Io {
  stat, _ := os.Stdin.Stat()
  if (stat.Mode() & os.ModeCharDevice) == 0 {
    return NewStdOut()
  } else {
    return NewFileIo(path)
  }
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}
