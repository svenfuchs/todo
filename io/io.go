package io

// var stdin *os.File

type Io interface {
  ReadLines() ([]string, error)
  WriteLines([]string) error
  AppendLines([]string) error
}

func NewIo(path string) Io {
  if path == "" {
    return StdIo{}
  } else {
    return FileIo{ path: path }
  }
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}
